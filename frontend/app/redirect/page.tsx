'use client';

import {useRouter} from "next/navigation";
import {useEffect} from "react";
import Connect from '../../components/Connect'
import { useWeb3Modal } from '@web3modal/wagmi/react'



export default function Redirect() {
    const router = useRouter();
    const { open } = useWeb3Modal()

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