'use client';

import {useRouter} from "next/navigation";
import {useEffect} from "react";
import Connect from '../../components/Connect'
import { useWeb3ModalProvider, useWeb3ModalAccount } from '@web3modal/ethers/react'
import { BrowserProvider, Contract, ethers, formatUnits } from 'ethers'

import { parseEther } from 'viem' 
// import { estimateGas } from "viem/_types/actions/public/estimateGas";


export default function Redirect() {
    const router = useRouter();
    const { address } = useWeb3ModalAccount()
   

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

    // useEffect(() => {
    //     async function openWallet() {
    //       // open()
          
  
    //     }

    //     openWallet()
    //     sendTransaction({to: `0x${"47dEAF612F0769d99aDB653bA2d22bba79F26C42"}`, value: parseEther("0.1")})
    // })
    async function submitTx() { console.log(address) }
    
 

      const handleCloseButtonClick = () => {
        // Close the current tab
        window.close();
      };


    return (
        <div>
            
            <Connect />
            <button onClick={submitTx}>Stake</button>
            <button onClick={handleCloseButtonClick}>Close Tab</button>
            {/* <button onClick={() => signMessage({ message: 'hello world' })}>Sign message</button> */}
            {/* <button onClick={() => sendTransaction( { to: `0x${"47dEAF612F0769d99aDB653bA2d22bba79F26C42"}`, value: parseEther("0.2") })}>Send Tx</button> */}
        </div>
    );
}



// exit page when done
// async function submitTx() {

//   const { walletProvider } = useWeb3ModalProvider()
//   if (!isConnected) throw Error('User disconnected')
//   if (!walletProvider) throw Error('Wallet Provider Abscent')
//   const ethersProvider = new BrowserProvider(walletProvider)
// console.log(ethersProvider)
  // const signer = await ethersProvider.getSigner()
  // let estimateGas = await ethersProvider.estimateGas({ to: `0x${"47dEAF612F0769d99aDB653bA2d22bba79F26C42"}`, value: parseEther("0.2")})

  // console.log(estimateGas)

  // await signer.sendTransaction({ to: `0x${"47dEAF612F0769d99aDB653bA2d22bba79F26C42"}`, value: parseEther("0.2")})
// }