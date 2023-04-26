// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "../common/AddressStorage.sol";
import "./DKG.sol";
import "./SignerOwnable.sol";
import "./ContractKeys.sol";
import "./ContractRegistry.sol";
import "./SlashingVoting.sol";
import "./ValidatorRewardDistributionPool.sol";

contract Staking is ContractKeys, SignerOwnable, Initializable {
    enum ValidatorStatus {
        INACTIVE,
        ACTIVE,
        SLASHED
    }

    struct ValidatorInfo {
        address validator;
        uint256 stake;
        ValidatorStatus status;
    }

    struct WithdrawalAnnouncement {
        uint256 amount;
        uint256 time;
    }

    uint256 public minimalStake;
    uint256 public withdrawalPeriod;
    uint256 public totalStake;
    mapping(address => ValidatorInfo) public stakes;
    mapping(address => WithdrawalAnnouncement) public withdrawalAnnouncements;
    ContractRegistry public contractRegistry;
    AddressStorage public addressStorage;

    event MinimalStakeUpdated(uint256 minimalStake);
    event WithdrawalPeriodUpdated(uint256 withdrawalPeriod);
    event ContractRegistryUpdated(address contractRegistry);

    modifier onlyNotSlashed() {
        require(stakes[msg.sender].status != ValidatorStatus.SLASHED, "Staking: validator is slashed");
        _;
    }

    modifier onlySlashingVoting() {
        require(msg.sender == address(_slashingVotingContract()), "Staking: not a slashing voting");
        _;
    }

    modifier onlyValidatorRewardDistributionPool() {
        require(
            msg.sender == address(_validatorRewardDistributionPoolContract()),
            "Staking: only ValidatorRewardDistributionPool contract"
        );
        _;
    }

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    function initialize(
        address _signerGetterAddress,
        uint256 _minimalStake,
        uint256 _withdrawalPeriod,
        address _contractRegistry,
        address _validatorStorage
    ) external initializer {
        _setSignerGetter(_signerGetterAddress);
        setMinimalStake(_minimalStake);
        setWithdrawalPeriod(_withdrawalPeriod);
        contractRegistry = ContractRegistry(_contractRegistry);
        addressStorage = AddressStorage(_validatorStorage);
    }

    function addRewardsToStake(address _validator, uint256 _amount) external onlyValidatorRewardDistributionPool {
        stakes[_validator].stake += _amount;
        totalStake += _amount;
    }

    function isValidatorActive(address _validator) external view returns (bool) {
        return (stakes[_validator].status == ValidatorStatus.ACTIVE);
    }

    function isValidatorSlashed(address _validator) external view returns (bool) {
        return stakes[_validator].status == ValidatorStatus.SLASHED;
    }

    function setMinimalStake(uint256 _minimalStake) public onlySigner {
        minimalStake = _minimalStake;
        emit MinimalStakeUpdated(_minimalStake);
    }

    function setWithdrawalPeriod(uint256 _withdrawalPeriod) public onlySigner {
        withdrawalPeriod = _withdrawalPeriod;
        emit WithdrawalPeriodUpdated(_withdrawalPeriod);
    }

    function slash(address _validator) public onlySlashingVoting {
        stakes[_validator].status = ValidatorStatus.SLASHED;
        _removeValidator(_validator);
    }

    function announceWithdrawal(uint256 _amount) public onlyNotSlashed {
        require(_amount <= stakes[msg.sender].stake, "Staking: amount must be <= to stake");
        withdrawalAnnouncements[msg.sender].amount = _amount;
        // solhint-disable-next-line not-rely-on-time
        withdrawalAnnouncements[msg.sender].time = block.timestamp;

        if (stakes[msg.sender].stake - _amount < minimalStake && addressStorage.contains(msg.sender)) {
            stakes[msg.sender].status = ValidatorStatus.INACTIVE;
            _removeValidator(msg.sender);
        }
    }

    function revokeWithdrawal() public onlyNotSlashed {
        require(withdrawalAnnouncements[msg.sender].amount > 0, "Staking: not announced");

        uint256 amount = withdrawalAnnouncements[msg.sender].amount;

        withdrawalAnnouncements[msg.sender].amount = 0;
        withdrawalAnnouncements[msg.sender].time = 0;

        if (
            stakes[msg.sender].status == ValidatorStatus.INACTIVE && amount + stakes[msg.sender].stake >= minimalStake
        ) {
            stakes[msg.sender].validator = msg.sender;
            stakes[msg.sender].status = ValidatorStatus.ACTIVE;
            _addValidator(msg.sender);
        }
    }

    function withdraw() public onlyNotSlashed {
        require(withdrawalAnnouncements[msg.sender].amount > 0, "Staking: amount must be greater than zero");
        require(
            // solhint-disable-next-line not-rely-on-time
            withdrawalAnnouncements[msg.sender].time + withdrawalPeriod <= block.timestamp,
            "Staking: withdrawal period not passed"
        );

        uint256 withdrawalAmount = withdrawalAnnouncements[msg.sender].amount;

        require(withdrawalAmount <= stakes[msg.sender].stake, "Staking: amount must be <= to validator stake");

        stakes[msg.sender].stake -= withdrawalAmount;
        totalStake -= withdrawalAmount;

        withdrawalAnnouncements[msg.sender].amount = 0;
        withdrawalAnnouncements[msg.sender].time = 0;

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = msg.sender.call{value: withdrawalAmount, gas: 21000}("");
        require(success, "Staking: transfer failed");
    }

    function stake() public payable onlyNotSlashed {
        require(msg.value > 0, "Staking: amount must be greater than zero");

        if (
            stakes[msg.sender].status == ValidatorStatus.INACTIVE &&
            msg.value + stakes[msg.sender].stake >= minimalStake
        ) {
            stakes[msg.sender].validator = msg.sender;
            stakes[msg.sender].status = ValidatorStatus.ACTIVE;
            _addValidator(msg.sender);
        }

        stakes[msg.sender].stake += msg.value;
        totalStake += msg.value;
    }

    function getValidators() public view returns (address[] memory) {
        return addressStorage.getAddresses();
    }

    function getValidatorsCount() public view returns (uint256) {
        return addressStorage.size();
    }

    function listValidators(uint256 _offset, uint256 _limit) public view returns (ValidatorInfo[] memory) {
        address[] memory validators = getValidators();
        ValidatorInfo[] memory result = new ValidatorInfo[](_limit);
        for (uint256 i = _offset; i < _offset + _limit; i++) {
            result[i - _offset] = stakes[validators[i]];
        }

        return result;
    }

    function getStake(address _validator) public view returns (uint256) {
        return stakes[_validator].stake;
    }

    function getTotalStake() public view returns (uint256) {
        return totalStake;
    }

    function _addValidator(address _validator) private {
        DKG dkg = _dkgContract();

        addressStorage.mustAdd(_validator);
        dkg.updateGeneration();
    }

    function _removeValidator(address _validator) private {
        DKG dkg = _dkgContract();

        addressStorage.mustRemove(_validator);
        dkg.updateGeneration();
    }

    function _dkgContract() private view returns (DKG) {
        return DKG(contractRegistry.getContract(DKG_KEY));
    }

    function _slashingVotingContract() private view returns (SlashingVoting) {
        return SlashingVoting(contractRegistry.getContract(SLASHING_VOTING_KEY));
    }

    function _validatorRewardDistributionPoolContract() private view returns (ValidatorRewardDistributionPool) {
        return
            ValidatorRewardDistributionPool(
                payable(contractRegistry.getContract(VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY))
            );
    }
}
