// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IERC20} from "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";


contract VersusIntegration {
  event EthDeposit(uint256 indexed fid, uint256 indexed amount, address indexed depositor);
  event VUSDCDeposit(uint256 indexed fid, uint256 indexed amount, address indexed depositor);

  address public vusdcAddress;

   constructor(address VersusUSDC) {
      vusdcAddress = VersusUSDC;
   }
  

   function depositVusdc(uint256 fid, uint256 amount) public payable {
      IERC20 vUSDC = IERC20(vusdcAddress);
      bool success = vUSDC.transferFrom(msg.sender, address(this), amount);
      require(success, "Transfer failed");
      emit VUSDCDeposit(fid, amount, msg.sender);
   }
   
   function withdrawVusdc(uint256 amount) public {
      IERC20 vUSDC = IERC20(vusdcAddress);
      bool success = vUSDC.transfer(msg.sender, amount);
      require(success, "Transfer failed");
   }


    /**
     * deposit Ether
     */
     function depositEth(uint256 fid) public payable {
        emit EthDeposit(fid, msg.value, msg.sender);
     }

    /**
    * withdraw
    */
     function withdrawEth(uint256 _amount) public {
        payable(msg.sender).transfer(_amount);
     }
}




// Deployer: 0x4211b8344e54C29993ba852CDe47B8dBA32936C8
// Deployed to: 0x3725db93a289Fdc9b2Fb9606a71952AB7cfbD14a
// Transaction hash: 0xb20aecb735c6565182271517869110b1c2c22c1a858b6d03f2580320be5cbe54
// depsoit
// cast send 0x3725db93a289Fdc9b2Fb9606a71952AB7cfbD14a "depositEth(uint256)" 245 --value 0.00003ether --rpc-url $BASE_TEST_RPC_URL --private-key $BASE_TEST_PRIV_KEY
// withdraw
// cast send 0x3725db93a289Fdc9b2Fb9606a71952AB7cfbD14a "withdrawEth(uint256)" 30000000000000 --rpc-url $BASE_TEST_RPC_URL --private-key $BASE_TEST_PRIV_KEY




// recent 
// Deployer: 0x4211b8344e54C29993ba852CDe47B8dBA32936C8
// Deployed to: 0xd9D454387F1cF48DB5b7D40C5De9d5bD9a92C1F8
// Transaction hash: 0x1badb03bec24fd2e5c25a98a07e86058f5e7d06e826f5ba5354b02858552da67

