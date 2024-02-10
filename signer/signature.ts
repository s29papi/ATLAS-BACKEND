import { mnemonicToAccount } from "viem/accounts";
import * as dotenv from 'dotenv';
dotenv.config();
/*** EIP-712 helper code ***/

const SIGNED_KEY_REQUEST_VALIDATOR_EIP_712_DOMAIN = {
  name: "Farcaster SignedKeyRequestValidator",
  version: "1",
  chainId: 10,
  verifyingContract: "0x00000000fc700472606ed4fa22623acf62c60553",
} as const;

const SIGNED_KEY_REQUEST_TYPE = [
  { name: "requestFid", type: "uint256" },
  { name: "key", type: "bytes" },
  { name: "deadline", type: "uint256" },
] as const;

const publicKey = process.env.PUBLIC_KEY; // Create and fetch this using the Neynar APIs
const mnemonic = process.env.MNEMONIC;


const appFid = process.env.FID; // Your app's fid
const accountFn = () => {
    let acc;
    if (mnemonic) {
        acc = mnemonicToAccount(mnemonic);
    }
    return acc // Your app's mnemonic
}

let account = accountFn();

const deadline = Math.floor((Date.now() + 7 * 24 * 60 * 60 * 1000) / 1000);

async function signAndLogSignature() {
    if (account && appFid && publicKey) {
            const signature = await account.signTypedData({
                domain: SIGNED_KEY_REQUEST_VALIDATOR_EIP_712_DOMAIN,
                types: {
                  SignedKeyRequest: SIGNED_KEY_REQUEST_TYPE,
                },
                primaryType: "SignedKeyRequest",
                message: {
                  requestFid: BigInt(appFid),
                  key: `0x${publicKey}`,
                  deadline: BigInt(deadline),
                },
              });
            
              console.log("Signature:", signature);
    }
  }
  
  signAndLogSignature();

  console.log(deadline)







