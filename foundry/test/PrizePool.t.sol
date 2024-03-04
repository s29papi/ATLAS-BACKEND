// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import {PrizePool} from "../src/PrizePool.sol"; 

contract PrizePoolTest is Test {
            PrizePool public prizepool;

            function setUp() public {
                prizepool = new PrizePool();
            }

            // Receive ETH from wallet
            receive() external payable {}

            function test_deposit_and_withdraw() public {
                address user1 = makeAddr("bixxy");
                hoax(user1, 1000000000000000000);
                bytes memory depositPayload = abi.encodeWithSignature("depositEth(uint256)", 234);
                (bool depositSuccess, ) = address(prizepool).call{value: 1000000000000000000}(depositPayload);
                require(depositSuccess);
                require(address(prizepool).balance == 1000000000000000000);
                vm.prank(user1);
                prizepool.withdrawEth(1000000000000000000);
                require(address(prizepool).balance == 0);
                console.log(address(this).balance);
                    
            }

            // function test_withdraw() public {
            //           // withdraw
            //           deal(address(prizepool), 1000000000000000000);
               
            //     // require(address(prizepool).balance == 0);            
            //     // require(address(1).balance == 1000000000000000000);  
            // }
}