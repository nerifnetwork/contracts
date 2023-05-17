// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract RegistryBalance {
    struct Balance {
        address user;
        uint256 balance;
    }

    mapping(address => uint256) private indexMap;
    Balance[] private balances;

    function getBalances() external view returns (Balance[] memory) {
        return balances;
    }

    function getBalance(address user) public view returns (uint256) {
        if (!_hasBalance(user)) {
            return 0;
        }

        return balances[indexMap[user] - 1].balance;
    }

    function _setBalance(address user, uint256 balance) internal {
        if (_hasBalance(user)) {
            balances[indexMap[user] - 1] = Balance(user, balance);
        } else {
            balances.push(Balance(user, balance));
            indexMap[user] = balances.length;
        }

        _checkBalanceEntry(user);
    }

    function _checkBalanceEntry(address user) private view {
        uint256 index = indexMap[user];
        assert(index <= balances.length);

        if (_hasBalance(user)) {
            assert(index > 0);
        } else {
            assert(index == 0);
        }
    }

    function _hasBalance(address user) private view returns (bool) {
        return indexMap[user] > 0;
    }
}
