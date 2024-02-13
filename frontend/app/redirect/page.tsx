'use client';
import {useRouter} from "next/navigation";
import {useEffect} from "react";

export default function Redirect() {
    const router = useRouter();

    useEffect(() => {
        const redirectUrl = '';

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
