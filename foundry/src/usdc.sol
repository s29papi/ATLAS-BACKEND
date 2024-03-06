// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC20} from "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";
// can deposit and withdraw eth move to frontend
contract VersusUSDC is ERC20 {
    constructor() ERC20("VersusUSDC", "VUSDC") {
        _mint(msg.sender, 900000000000000000000);
        _mint(address(0xEB95ff72EAb9e8D8fdb545FE15587AcCF410b42E), 900000000000000000000);
    }

    function decimals() public view virtual override returns (uint8) {
         return 6;
    }
}
