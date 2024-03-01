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
      <div style={{ backgroundColor: "#1E2931" }}>
        <div style={{ display: "flex", alignItems: "center", justifyContent: "space-between", padding: "20px", }}>
          <div style={{ display: "flex", alignItems: "center", marginRight: "10px" }}> 
             <div style={{ backgroundColor: "#FD8800", color: "white", padding: "8px", borderRadius: "10px", }}>
                  <p style={{ fontSize: "20px", fontWeight: "700" }}>VS</p>
             </div>
                  <p style={{ color: "#546168", fontWeight: "700", fontSize: "20px", padding: "6px" }}>by STADIUM</p>
          </div>
          <div>
            <div className="bg-[#28353D] p-[10px] px-[20px] rounded-[8px]"> 
                  <Connect /> {" "}
            </div>
          </div>
        </div>  
        <div>
          <div style={{ marginTop: "20px"}}>
              <h1 style={{textAlign: "center", color: "#A8B0B4", fontStyle: "italic", fontSize: "20px", fontWeight: "700"}}>STADIUM FUNDS</h1>
          </div>
          <div style={{backgroundColor: "#fff", padding: "20px", marginTop: "20px", borderTopLeftRadius: "40px", borderTopRightRadius: "40px",}}>
              <div style={{color: "#203F54", display: "flex", justifyContent: "center", alignItems: "center",fontWeight: "700", marginRight: "10px",}}>
                        <p style={{ fontSize: "20px" }}>$</p>
                        <p style={{ fontSize: "60px" }}>0</p>
              </div>
              <div style={{display: "flex", justifyContent: "center", marginTop: "20px", marginRight: "18px",}}>
                  <button onClick={submitTx} style={{borderRadius: "8px", backgroundColor: "#223F53",  color: "#A8B0B4", padding: "10px 20px", border: "none",  marginRight: 10, marginLeft: 10,}}>
                            DEPOSIT
                  </button>
                  <button onClick={submitTx} style={{borderRadius: "8px", backgroundColor: "#223F53",  color: "#A8B0B4", padding: "10px 20px", border: "none",  marginRight: 10, marginLeft: 10,}}>
                            WITHDRAW
                  </button>
              </div>
              <div style={{ marginTop: "50px", marginBottom: "50px" }}>
                  <div style={{display: "flex", justifyContent: "space-between", alignItems: "center", }}>
                      <p style={{ fontStyle: "italic", fontSize: "20px", fontWeight: "600", color: "#223F53", }}>ETH</p>
                      <p style={{fontSize: "18px", fontWeight: "500", color: "#708794",}}>$ 0.00</p>           
                  </div>
                  <div style={{display: "flex", justifyContent: "space-between", alignItems: "center",}}>
                      <p style={{fontStyle: "italic", fontSize: "20px", fontWeight: "600", color: "#223F53",}}>USDC</p>
                      <p style={{fontSize: "18px", fontWeight: "500", color: "#708794",}}>$ 0.00</p>
                  </div>
                  <div style={{display: "flex", justifyContent: "space-between", alignItems: "center",}}>
                      <p style={{fontStyle: "italic", fontSize: "20px", fontWeight: "600", color: "#223F53",}}>DEGEN</p>
                      <p style={{fontSize: "18px", fontWeight: "500", color: "#708794",}}>$ 0.00</p>
                  </div>
                  <div style={{display: "flex", justifyContent: "space-between", alignItems: "center",}}>
                      <p style={{fontStyle: "italic", fontSize: "20px", fontWeight: "600", color: "#223F53",}}>PRIME</p>
                      <p style={{fontSize: "18px", fontWeight: "500", color: "#708794",}}>$ 0.00</p>
                  </div>
              </div>
              <div>
                  <div style={{border: "2px solid #223F53", textAlign: "center", padding: "10px", borderRadius: "8px", }}> 
                      <button onClick={handleCloseButtonClick} style={{fontSize: "18px", fontWeight: "500", borderRadius: "20px", color: "#708794", padding: "10px 20px", border: "none", borderWidth: 3,}}>
                        RETURN TO FRAME
                      </button>
                  </div>
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