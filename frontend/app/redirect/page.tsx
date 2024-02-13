'use client';
import {useRouter} from "next/navigation";
import {useEffect} from "react";

export default function RedirectPage() {
    const router = useRouter();

    useEffect(() => {
        const redirectUrl = 'https://twitter.com';

        const width = 600;
        const height = 400;
        const left = (window.innerWidth - width) / 2;
        const top = (window.innerHeight - height) / 2;
        const options = `width=${width},height=${height},left=${left},top=${top}`;
    

        // Perform the redirect
        // window.location.href = redirectUrl; // For a full page reload redirect
        window.open(redirectUrl)
        // Or use Next.js router for client-side redirect (comment out the line above if using this)
        // router.push(youtubeUrl);
    }, [router]);

    return (
        <div>
            <p>Redirecting...</p>
        </div>
    );
}



// 'use client';
// // import {useRouter} from "next/navigation";
// // import { useSearchParams } from 'next/navigation'
// // import {useEffect} from "react";


// export default function RedirectPage() {
//     const router = useRouter();
//     const searchParams = useSearchParams()
//     let dAddr = searchParams.get("DestinationAddress")
//     let dchainId = searchParams.get("DestinationChainId")
//     let schainId = searchParams.get("SourceChainId")
//     let iamount = searchParams.get("InputAmount")
//     let userInfo = "DestinationAddress=" + dAddr + "&" + "DestinationChainId=" +  dchainId + "&" + "SourceChainId=" + schainId + "&" + "InputAmount=" + iamount; 
 
    
  

//     useEffect(() => {
//         const redirectUrl = 'https://socket-pay.vercel.app/redirect/signTx' + '?' + `${userInfo}`;

//         // Perform the redirect
//         window.location.href = redirectUrl; // For a full page reload redirect
//         // Or use Next.js router for client-side redirect (comment out the line above if using this)
//         // router.push(youtubeUrl);
//     }, [router]);

//     return (
//         <div>
//             <p>Redirecting...</p>
//         </div>
//     );
// }  