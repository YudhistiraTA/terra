import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export const config = {
  matcher: [
    /*
     * Match all request paths except for the ones starting with:
     * - api (API routes)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico, sitemap.xml, robots.txt (metadata files)
     */
    "/((?!api|_next/static|_next/image|favicon.ico|sitemap.xml|robots.txt|login).*)",
  ],
};

function redirectToLogin(request: NextRequest) {
  const response = NextResponse.redirect(new URL("/login", request.nextUrl));
  response.cookies.delete("token");
  return response;
}

export function middleware(request: NextRequest) {
  const token = request.cookies.get("token");
  if (!token) {
    return redirectToLogin(request);
  }
  const expiry = new Date(token.value.split(".")[1]);
  if (expiry < new Date()) {
    return redirectToLogin(request);
  }
}
