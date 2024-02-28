'use client';

import {useRouter} from "next/navigation";
import {useEffect} from "react";
import Connect from '../../../components/Connect';
import { useWeb3ModalProvider, useWeb3ModalAccount } from '@web3modal/ethers/react';
import { BrowserProvider, Contract, ethers, formatUnits } from 'ethers';
import { parseEther } from 'viem';

type Props = {
  params: { gameId: string }
  searchParams: { [key: string]: string | string[] | undefined }
}


export default async function StakePage({ params, searchParams }: Props) {
    const router = useRouter();
    const { walletProvider } = useWeb3ModalProvider()
    let fid = searchParams["fid"];

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
      console.log(fid)
      // if (!walletProvider) throw Error('Wallet Provider Abscent')
      // const ethersProvider = new BrowserProvider(walletProvider)
      // const signer = await ethersProvider.getSigner()
      // let estimateGas = await ethersProvider.estimateGas({ to: `0x${"47dEAF612F0769d99aDB653bA2d22bba79F26C42"}`, value: parseEther("0.2")})
      // let sentTx = await signer.sendTransaction({ to: `0x${"47dEAF612F0769d99aDB653bA2d22bba79F26C42"}`, value: parseEther("0.2"), gasLimit: estimateGas})
      // let resolvedTx = await sentTx.wait()
      // // we keep id in the cookie
      // // db stores id and tx hash
      // console.log(resolvedTx?.hash)
      
     }
    
 

      const handleCloseButtonClick = () => {
        // Close the current tab
        window.close();
      };


    return (
      <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', height: '100vh' }}>
        Start By Clicking Connect To Connect your Wallet.
        <div style={{ display: 'flex', gap: '20px', marginTop: '20px' }}>
          <Connect />
            <button onClick={submitTx} style={{ borderRadius: '20px', backgroundColor: 'rgb(51, 204, 153)', color: 'white', padding: '10px 20px', border: 'none' }}>Stake</button>
           <button onClick={handleCloseButtonClick} style={{ borderRadius: '20px', backgroundColor: 'red', color: 'white', padding: '10px 20px', border: 'none' }}>Finish</button>
        </div>
    </div>
    );
}

