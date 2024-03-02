'use client';

import {useRouter} from "next/navigation";
import {useEffect} from "react";
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


export default function StakePage({ params, searchParams }: Props) {
    const router = useRouter();
    const { walletProvider } = useWeb3ModalProvider()
    let paramsFid = searchParams["fid"];
    let fid: string;
    if (paramsFid) {
      fid = paramsFid.toString()
    }
    // let fid = searchParams["fid"];
   

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


    async function submitTx() { 
      if (!walletProvider) throw Error('Wallet Provider Abscent')
      const ethersProvider = new BrowserProvider(walletProvider)
      const signer = await ethersProvider.getSigner()
      let estimateGas = await ethersProvider.estimateGas({
         to: `0x${"47dEAF612F0769d99aDB653bA2d22bba79F26C42"}`,
         value: parseEther("0.2"), 
         data: fid
        });
      let sentTx = await signer.sendTransaction({
         to: `0x${"47dEAF612F0769d99aDB653bA2d22bba79F26C42"}`, 
         value: parseEther("0.2"), 
         data: fid,
         gasLimit: estimateGas
        });
      let resolvedTx = await sentTx.wait()
      // we keep id in the cookie
      // db stores id and tx hash
      console.log(resolvedTx?.hash)
      
     }
    
 

      const handleCloseButtonClick = () => {
        // Close the current tab
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

      <div className="flex h-[37.69px] items-center justify-center rounded-[10px] bg-[#66757F] px-4 font-semibold text-white opacity-[70%]">
        <p>Connect</p>
      </div>
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
          <div onClick={submitTx} className="cursor-pointer rounded-[10px] bg-[#223F53] px-4 py-2 font-semibold hover:bg-[#213D52] hover:opacity-[50%]">Deposit</div>
          <div onClick={submitTx} className="cursor-pointer rounded-[10px] bg-[#223F53] px-4 py-2 font-semibold hover:bg-[#213D52] hover:opacity-[50%]">Withdraw</div>
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

      <div onClick={handleCloseButtonClick} className="mx-[24px] cursor-pointer hover:bg-[#223F53] hover:text-white mb-[16px] mt-[38px] rounded-[8px] border-[2px] border-[#223F53] text-center">
        <p className="text-[18px] font-bold">RETURN TO FRAME</p>
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




{/* <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', height: '100vh' }}>
Start By Clicking Connect To Connect your Wallet..
<div style={{ display: 'flex', gap: '20px', marginTop: '20px' }}>
  <Connect />
   
   
</div>
</div> */}


{/*      <div style={{ backgroundColor: '#1E2931' }}>
          <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', padding: '20px' }}>
            <div style={{ display: 'flex', alignItems: 'center', marginRight: '10px' }}>
              <div style={{ backgroundColor: '#FD8800', color: 'white', padding: '8px', borderRadius: '10px' }}>
                <p style={{ fontSize: '20px', fontWeight: '700' }}>VS</p>
              </div>
              <p style={{ color: '#546168', fontWeight: '700', fontSize: '20px' }}>by STADIUM</p>
            </div>

            <div><Connect /></div>
          </div>

          <div>
            <div style={{ marginTop: '20px' }}>
              <h1 style={{ textAlign: 'center', color: '#A8B0B4', fontStyle: 'italic', fontSize: '20px', fontWeight: '700' }}>STADIUM FUNDS</h1>
            </div>

            <div style={{ backgroundColor: '#fff', padding: '20px', marginTop: '20px', borderTopLeftRadius: '40px', borderTopRightRadius: '40px' }}>
              <div style={{ color: '#203F54', display: 'flex', justifyContent: 'center', alignItems: 'center', fontWeight: '700', marginRight: '10px' }}>
                <p style={{ fontSize: '20px' }}>$</p>
                <p style={{ fontSize: '60px' }}>0</p>
              </div>

              <div style={{ display: 'flex', justifyContent: 'center', marginTop: '20px', marginRight: '20px' }}>
                  <button onClick={submitTx} style={{ borderRadius: '8px', backgroundColor: '#223F53', color: '#A8B0B4', padding: '10px 20px', border: 'none' }}>DEPOSIT</button>
                  <button onClick={submitTx} style={{ borderRadius: '8px', backgroundColor: '#223F53', color: '#A8B0B4', padding: '10px 20px', border: 'none' }}>WITHDRAW</button>
              </div>

              <div style={{ marginTop: '50px', marginBottom: '50px' }}>
                <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <p style={{ fontStyle: 'italic', fontSize: '20px', fontWeight: '600', color: '#223F53' }}>ETH</p>
                  <p style={{ fontSize: '18px', fontWeight: '500', color: '#708794' }}>$ 0.00</p>
                </div>
                <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <p style={{ fontStyle: 'italic', fontSize: '20px', fontWeight: '600', color: '#223F53' }}>USDC</p>
                  <p style={{ fontSize: '18px', fontWeight: '500', color: '#708794' }}>$ 0.00</p>
                </div>
                <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <p style={{ fontStyle: 'italic', fontSize: '20px', fontWeight: '600', color: '#223F53' }}>DEGEN</p>
                  <p style={{ fontSize: '18px', fontWeight: '500', color: '#708794' }}>$ 0.00</p>
                </div>
                <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <p style={{ fontStyle: 'italic', fontSize: '20px', fontWeight: '600', color: '#223F53' }}>PRIME</p>
                  <p style={{ fontSize: '18px', fontWeight: '500', color: '#708794' }}>$ 0.00</p>
                </div>
              </div>

              <div>
                <div style={{ border: '2px solid #223F53', textAlign: 'center', padding: '10px', borderRadius: '8px' }}>
                  <button onClick={handleCloseButtonClick} style={{ fontSize: '18px', fontWeight: '500', borderRadius: '20px', backgroundColor: 'red', color: '#708794', padding: '10px 20px', border: 'none' }}>RETURN TO FRAME</button>
                </div>

                <div style={{ width: '100%', height: '200px', backgroundRepeat: 'no-repeat', backgroundSize: 'cover', marginTop: '50px', borderRadius: '12px', position: 'relative', backgroundImage: `url(${Image})`, backgroundPosition: 'center' }}>
                  <div style={{ position: 'absolute', top: '0', left: '0', width: '100%', height: '100%', borderRadius: '12px', padding: '20px', display: 'flex', flexDirection: 'column', justifyContent: 'space-between', backgroundImage: 'linear-gradient(to bottom, rgba(0, 0, 0, 0), rgba(0, 0, 0, 0.8))' }}>
                    <p style={{ width: '189px', color: 'white', fontStyle: 'italic', fontWeight: '600' }}>PROVE YOUR SKILL GET YOUR SKIN IN THE GAME</p>

                    <div style={{ color: 'white', fontSize: '12px' }}>
                      <p>🔎 Create or find a challenge</p>
                      <p>💰 Stake your tokens</p>
                      <p>🎮 Compete in the challenge</p>
                      <p>🏆 Earn tokens by winning</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div style={{ backgroundColor: '#1F2A32', width: '100%', color: 'white', padding: '8px', paddingTop: '16px', paddingBottom: '16px', display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
            <p>Follow @versus | Support @sirsu + @hidd3n Powered by Base, BLVKHVND, Stadium</p>
            <img src={FooterIcon} alt="" />
          </div>
</div>*/}