'use client';

import {useRouter} from "next/navigation";
import {useEffect} from "react";
import Connect from '../../components/Connect'
// import { useWeb3Modal } from '@web3modal/wagmi/react'
import { useSendTransaction, useSignMessage, usePrepareTransactionRequest  } from 'wagmi' 
import { prepareTransactionRequest } from '@wagmi/core'
import { parseEther } from 'viem' 
import { config } from "../../config"



export default function Redirect() {
    const router = useRouter();
    // const { open } = useWeb3Modal()
    const { data: hash, sendTransaction } = useSendTransaction({config,}) 
    const { signMessage } = useSignMessage()
    const result = usePrepareTransactionRequest({
      to: '0x70997970c51812dc3a010c7d01b50e0d17dc79c8',
      value: parseEther('1'),
    })

    // const { data, isSuccess, sendTransaction, error } = useSendTransaction(result);
   

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

    async function submitTx() {
     let request = await  prepareTransactionRequest(config, {
          to: '0x70997970c51812dc3a010c7d01b50e0d17dc79c8',
          value: parseEther('1'),
        })
      // usePrepareTransactionRequest(request)
    }

      const handleCloseButtonClick = () => {
        // Close the current tab
        window.close();
      };


    return (
        <div>
            
            <Connect />
            <button onClick={async () => { await submitTx() }}>Stake</button>
            <button onClick={handleCloseButtonClick}>Close Tab</button>
            <button onClick={() => signMessage({ message: 'hello world' })}>Sign message</button>
            <button onClick={() => sendTransaction( { to: `0x${"47dEAF612F0769d99aDB653bA2d22bba79F26C42"}`, value: parseEther("0.2") })}>Send Tx</button>
        </div>
    );
}



// exit page when done