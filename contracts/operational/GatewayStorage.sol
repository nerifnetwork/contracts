// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "../interfaces/IGateway.sol";

// GatewayStorage is the gateways storage contract.
contract GatewayStorage is Ownable, Initializable {
    // Gateway is the gateway model
    struct Gateway {
        IGateway gateway;
        address owner;
    }

    address public registry;
    mapping(address => uint256) internal indexMap;
    Gateway[] internal gateways;

    // onlyRegistry permits transactions coming from the registry contract.
    modifier onlyRegistry() {
        require(registry == msg.sender, "GatewayStorage: operation is not permitted");
        _;
    }

    function initialize(address _registry, Gateway[] memory _gateways) external virtual initializer {
        registry = _registry;

        for (uint256 i = 0; i < _gateways.length; i++) {
            gateways.push(_gateways[i]);
            indexMap[_gateways[i].owner] = gateways.length;
        }
    }

    function setRegistry(address _registry) external onlyOwner {
        registry = _registry;
    }

    function mustSet(address owner, address gateway) external onlyRegistry {
        _set(owner, IGateway(gateway));
    }

    function mustRemove(address owner) external onlyRegistry {
        require(_remove(owner), "GatewayStorage: failed to remove gateway");
    }

    function clear() external onlyOwner returns (bool) {
        for (uint256 i = 0; i < gateways.length; i++) {
            delete indexMap[gateways[i].owner];
        }

        delete gateways;

        return true;
    }

    function size() external view returns (uint256) {
        return gateways.length;
    }

    function getGateways() external view returns (Gateway[] memory) {
        return gateways;
    }

    function getGateway(address owner) external view returns (IGateway) {
        if (!contains(owner)) {
            return IGateway(address(0x0));
        }

        return gateways[indexMap[owner] - 1].gateway;
    }

    function contains(address owner) public view returns (bool) {
        return indexMap[owner] > 0;
    }

    function _set(address owner, IGateway gateway) private onlyRegistry {
        if (contains(owner)) {
            gateways[indexMap[owner] - 1] = Gateway(gateway, owner);
        } else {
            gateways.push(Gateway(gateway, owner));
            indexMap[owner] = gateways.length;
        }

        _checkEntry(owner);
    }

    function _remove(address owner) private onlyRegistry returns (bool) {
        if (!contains(owner)) {
            return false;
        }

        uint256 index = indexMap[owner];

        uint256 lastListIndex = gateways.length - 1;
        Gateway memory lastListGateway = gateways[lastListIndex];
        if (lastListIndex != index - 1) {
            indexMap[lastListGateway.owner] = index;
            gateways[index - 1] = lastListGateway;
        }

        gateways.pop();
        delete indexMap[owner];

        _checkEntry(owner);

        if (lastListGateway.owner != owner) {
            _checkEntry(lastListGateway.owner);
        }

        return true;
    }

    function _checkEntry(address owner) private view {
        uint256 index = indexMap[owner];
        assert(index <= gateways.length);

        if (contains(owner)) {
            assert(index > 0);
        } else {
            assert(index == 0);
        }
    }
}
