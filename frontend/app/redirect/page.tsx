'use client';
import {useRouter} from "next/navigation";
import {useEffect} from "react";

// declare global {
//     interface Window {
//       ethereum?: any;
//     }
//   } 

export default function Redirect() {
    const router = useRouter();

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
          <w3m-button />
            // const [account] = await window.ethereum.request({
            //     method: 'eth_requestAccounts',
            //   });
            // return account;
        }

        openWallet()
    })

      const handleCloseButtonClick = () => {
        // Close the current tab
        window.close();
      };


    return (
        <div>
            <p>Redirecting...</p>
            <button onClick={handleCloseButtonClick}>Close Tab</button>
        </div>
    );
}



