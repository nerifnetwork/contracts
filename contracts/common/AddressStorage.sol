// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract AddressStorage is Ownable, Initializable {
    mapping(address => uint256) internal indexMap;
    address[] internal addrList;

    function initialize(address[] memory _addrList) external virtual initializer {
        for (uint256 i = 0; i < _addrList.length; i++) {
            addrList.push(_addrList[i]);
            indexMap[_addrList[i]] = addrList.length;
        }
    }

    function mustAdd(address _addr) external {
        require(_add(_addr), "AddressStorage: failed to add address");
    }

    function mustRemove(address _addr) external {
        require(_remove(_addr), "AddressStorage: failed to remove address");
    }

    function clear() external onlyOwner returns (bool) {
        for (uint256 i = 0; i < addrList.length; i++) {
            delete indexMap[addrList[i]];
        }

        delete addrList;

        return true;
    }

    function size() external view returns (uint256) {
        return addrList.length;
    }

    function getAddresses() external view returns (address[] memory) {
        return addrList;
    }

    function contains(address _addr) public view returns (bool) {
        return indexMap[_addr] > 0;
    }

    function _add(address _addr) private onlyOwner returns (bool) {
        if (contains(_addr)) {
            return false;
        }

        addrList.push(_addr);
        indexMap[_addr] = addrList.length;

        _checkEntry(_addr);

        return true;
    }

    function _remove(address _addr) private onlyOwner returns (bool) {
        if (!contains(_addr)) {
            return false;
        }

        uint256 id = indexMap[_addr];

        uint256 lastListID = addrList.length - 1;
        address lastListAddress = addrList[lastListID];
        if (lastListID != id - 1) {
            indexMap[lastListAddress] = id;
            addrList[id - 1] = lastListAddress;
        }

        addrList.pop();
        delete indexMap[_addr];

        _checkEntry(_addr);

        if (lastListAddress != _addr) {
            _checkEntry(lastListAddress);
        }

        return true;
    }

    function _checkEntry(address _addr) private view {
        uint256 id = indexMap[_addr];
        assert(id <= addrList.length);

        if (contains(_addr)) {
            assert(id > 0);
        } else {
            assert(id == 0);
        }
    }
}
