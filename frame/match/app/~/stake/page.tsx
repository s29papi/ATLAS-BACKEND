'use client';

import {useRouter} from "next/navigation";
import {useState, useEffect} from "react";
import Connect from '../../../components/Connect';
import { useWeb3ModalProvider, useWeb3ModalAccount } from '@web3modal/ethers/react';
import { BrowserProvider, Contract, ethers, formatUnits } from 'ethers';
import { parseEther } from 'viem';
import Image from "./assets/image.svg";
import FooterIcon from "./assets/footer-icon.svg";


type Props = {
  params: { gameId: string }
  searchParams: { [key: string]: string | string[] | undefined }
}
import PrizePool from "./contracts/PrizePool.json"
import IERC20 from "./contracts/IERC20.json"

export default function StakePage({ params, searchParams }: Props) {
    let prizePoolAddr = "0xd9D454387F1cF48DB5b7D40C5De9d5bD9a92C1F8";
    let vusdcAddr = "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913";
    const router = useRouter();
    const { walletProvider } = useWeb3ModalProvider()
    let paramsFid = searchParams["fid"];
    let fid: string;
    if (paramsFid) {
      fid = paramsFid.toString()
    }
    if (!paramsFid) return (<div>Resource does not exist</div>);
    const [newDeposit, setNewDeposit] = useState(false);
    const [depositData, setDepositData] = useState("");



    useEffect(() => {
        const handleBeforeUnload = (event: BeforeUnloadEvent) => {
          // Cancel the default behavior of closing the tab
          event.preventDefault();
          // Chrome requires the returnValue to be set
          event.returnValue = '';
        };
    
        // Add event listener to beforeunload event
        window.addEventListener('beforeunload', handleBeforeUnload);
    
        return () => {
          // Remove the event listener when component unmounts
          window.removeEventListener('beforeunload', handleBeforeUnload);
        };
    }, []);


   
    async function submitDepositTx() { 
      let amount = 1000000;
            // let fid_string = ethers.hexlify(ethers.toUtf8Bytes(fid));
      if (!walletProvider) throw Error('Wallet Provider Abscent')
      const ethersProvider = new BrowserProvider(walletProvider)
      const signer = await ethersProvider.getSigner()
      let versusUsdc = new ethers.Contract(vusdcAddr, IERC20.abi, signer)
      const txApprove = await versusUsdc.approve(prizePoolAddr, amount)
      await txApprove.wait();
      console.log(`Tx successful with hash: ${txApprove.hash}`);
      let stadiumPrizePool = new ethers.Contract(prizePoolAddr, PrizePool.abi, signer)
      const tx = await stadiumPrizePool.depositVusdc(fid, 1000000);
      await tx.wait();
      console.log(`Tx successful with hash: ${tx.hash}`);
     let jsonData = JSON.stringify({ "fid": parseInt(fid), "address": signer.address, "txhash": tx.hash, "amount": amount });
      setDepositData(jsonData)
      setNewDeposit(true);
    }
    useEffect(() => {
      if (newDeposit) {
        fetch("http://localhost:8080/api/worker/deposit-request", {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MDk2MDc1MjV9.QQlkhaYpdeFhpYV7W2AP2hxdCfy4o3KkScB-ODxtnAc'
            },
            body: depositData
        }).then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            console.log('POST request successful');
          }).catch(error => {
              console.error('There was a problem with your fetch operation:', error);
          });
            setNewDeposit(false);
      }
  }, [newDeposit, depositData]);
    
    async function submitWithdrawTx() { 
      // let fid_string = ethers.hexlify(ethers.toUtf8Bytes(fid));

      if (!walletProvider) throw Error('Wallet Provider Abscent')
      const ethersProvider = new BrowserProvider(walletProvider)
      const signer = await ethersProvider.getSigner()
      let stadiumPrizePool = new ethers.Contract(prizePoolAddr, PrizePool.abi, signer)
      const tx = await stadiumPrizePool.withdrawVusdc(2000000);
      await tx.wait();
    
      console.log(`Tx successful with hash: ${tx.hash}`);
     }


    
 

      const handleCloseButtonClick = () => {
        window.close();
      };


    return (
        <div className="bg-[#2D3941]">
        <div className="mx-auto text-[#213D52] md:w-[660px] overflow-hidden">
          <div className="bg-[#1F2A32]">
            <div className="flex justify-between pb-[35px] pl-[24px] pr-[16px] pt-[18px]">
              <div className="flex items-center gap-2">
                <p className="rounded-[10px] bg-[#FF8F00] p-2 font-bold text-[white] shadow-xl">VS</p>
                <span className="text-[21px] font-bold text-[#66757F]">by STADIUM</span>
              </div>

                <Connect />{""} 
            
            </div>

            <p className="mb-[25px] text-center text-[28px] italic text-white">STADIUM FUNDS</p>

            <div className="w-full rounded-t-[36px] bg-white pt-[39px] shadow-lg">

              <div className="flex justify-center">
                <div className="flex items-center gap-2">
                  <p className="text-[24px]">$</p>
                  <p className="text-[60px]">0</p>
                </div>
              </div>

              <div className="mt-[26px] flex justify-center">
                <div className="flex items-center gap-2 text-white">
                  <button onClick={submitDepositTx} className="cursor-pointer rounded-[10px] bg-[#223F53] px-4 py-2 font-semibold hover:bg-[#213D52] hover:opacity-[50%]">Deposit</button>
                  <button onClick={submitWithdrawTx} className="cursor-pointer rounded-[10px] bg-[#223F53] px-4 py-2 font-semibold hover:bg-[#213D52] hover:opacity-[50%]">Withdraw</button>
                </div>
              </div>

              <div className="mt-[40px] space-y-[15px] px-[25px] font-semibold">
                <div className="flex justify-between text-[18px]">
                  <p className="">DEPOSITED FUNDS</p>
                  <p>$0.00</p>
                </div>
                <div className="space-y-[12px] font-medium">

                  <div className="flex items-center justify-between text-[16px]">
                    <div>
                      <p className="">Base ETH</p>
                      <div className="flex gap-[2px] text-[12px] italic text-[##66757F]">
                        <span>0</span>
                        <p className="">ETH</p>
                      </div>
                    </div>
                    <p className="text-[18px] text-[#53697A]">$0.00</p>
                  </div>

                  <div className="flex items-center justify-between text-[16px]">
                    <div>
                      <p className="">USDC</p>
                      <div className="flex gap-[2px] text-[12px] italic text-[##66757F]">
                        <span>0</span>
                        <p className="">USDC</p>
                      </div>
                    </div>
                    <p className="text-[18px] text-[#53697A]">$0.00</p>
                  </div>

                  <div className="flex items-center justify-between text-[16px]">
                    <div>
                      <p className="">Degen (base)</p>
                      <div className="flex gap-[2px] text-[12px] italic text-[##66757F]">
                        <span>0</span>
                        <p className="">DEGEN</p>
                      </div>
                    </div>
                    <p className="text-[18px] text-[#53697A]">$0.00</p>
                  </div>

                  <div className="flex items-center justify-between text-[16px]">
                    <div>
                      <p className="">Prime (base)</p>
                      <div className="flex gap-[2px] text-[12px] italic">
                        <span>0</span>
                        <p className="text-[#66757F]">PRIME</p>
                      </div>
                    </div>
                    <p className="text-[18px] text-[#53697A]">$0.00</p>
                  </div>
                </div>
              </div>

              <div className="mx-[24px] cursor-pointer hover:bg-[#223F53] hover:text-white mb-[16px] mt-[38px] rounded-[8px] border-[2px] border-[#223F53] text-center">
                <button onClick={handleCloseButtonClick} className="text-[18px] font-bold">RETURN TO FRAME</button>
              </div>

              <div className="mx-[12px] rounded-[8px] pb-[16px]">
                <img className="w-full bg-cover" src="https://imagedelivery.net/BXluQx4ige9GuW0Ia56BHw/c663b05b-d3fd-472a-28b8-7ecf96926800/original" alt="" />
              </div>
            </div>

            <div className="px-[9px] pb-[12px] pt-[16px] text-white">
              <p>Follow <span className="underline">@versus</span> | Support <span className="underline">@sirsu</span>+<span className="underline">@hidd3n</span></p>
              <p>Powered by Base, BLVKHVND, Stadium</p>
            </div>
          </div>
        </div>
        </div>
    );
}



      // curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MDk2MDc1MjV9.QQlkhaYpdeFhpYV7W2AP2hxdCfy4o3KkScB-ODxtnAc" http://localhost:8080/api/worker/deposit-request --data '{ "fid": 3, "address": "0x67123a", "txhash": "0x3773828", "amount": 930}'
