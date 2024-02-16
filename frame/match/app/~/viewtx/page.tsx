'use client';
import {useRouter} from "next/navigation";
// import { useSearchParams } from 'next/navigation'
import {useEffect} from "react";


export default function RedirectPage() {
    const router = useRouter();
 

    useEffect(() => {
        const redirectUrl = 'https://basescan.org/tx/0x7337cf1577fc77d27f1c0461147348b1290725188ad370fdfa04c908fd17cf4d';

        // Perform the redirect
        window.location.href = redirectUrl; // For a full page reload redirect
        // Or use Next.js router for client-side redirect (comment out the line above if using this)
        // router.push(youtubeUrl);
    }, [router]);

    return (
        <div>
            <p>Redirecting...</p>
        </div>
    );
}    