'use client';
import {useRouter} from "next/navigation";
import {useEffect} from "react";

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



    // useEffect(() => {
    //     const redirectUrl = '';

    //     // Perform the redirect
    //     window.location.href = redirectUrl; // For a full page reload redirect
    //     // Or use Next.js router for client-side redirect (comment out the line above if using this)
    //     // router.push(youtubeUrl);
    // }, [router]);