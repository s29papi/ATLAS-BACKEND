// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;


// can deposit and withdraw eth move to frontend
contract PrizePool {
    event EthDeposit(uint256 indexed fid, uint256 indexed amount, address indexed depositor);
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