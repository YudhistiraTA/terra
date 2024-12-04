import assert from "assert";
import { enc, HmacSHA512, SHA256 } from "crypto-js";

export default function generateSignature({
  endpoint,
  method,
  body,
  token,
}: {
  method: string;
  endpoint: string;
  body?: string;
  token?: string;
}) {
  const hmacSecret = process.env.NEXT_PUBLIC_HMAC_SECRET;
  assert(hmacSecret, "NEXT_PUBLIC_HMAC_SECRET is not set");
  const timestamp = new Date().toISOString();

  let payload = `${method}:${endpoint}:${timestamp}`;
  if (token) {
    payload += `:${token}`;
  }
  if (body) {
    const bodyHash = SHA256(body);
    const bodyHex = enc.Hex.stringify(bodyHash);
    payload += `:${bodyHex}`;
  }

  const signatureHash = HmacSHA512(payload, hmacSecret);
  const signature = enc.Base64.stringify(signatureHash);

  return {
    signature,
    timestamp,
  };
}
