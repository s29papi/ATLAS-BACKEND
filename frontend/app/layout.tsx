import './globals.css'

import { Web3ModalProvider } from '../context/Web3Modal'

export const metadata = {
  title: 'Web3Modal',
  description: 'Web3Modal Example'
}


export const viewport = {
  width: 'device-width',
  initialScale: 1.0,
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body><Web3ModalProvider>{children}</Web3ModalProvider></body>
    </html>
  );
}
