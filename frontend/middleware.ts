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
    "/((?!api|_next/static|_next/image|favicon.ico|sitemap.xml|robots.txt).*)",
  ],
};

function redirectToLogin(request: NextRequest) {
  const response = NextResponse.redirect(new URL("/login", request.nextUrl));
  response.cookies.delete("sessionToken");
  response.cookies.delete("refreshToken");
  return response;
}

export function middleware(request: NextRequest) {
  const sessionToken = request.cookies.get("sessionToken");
  const refreshToken = request.cookies.get("refreshToken");
  if (request.nextUrl.pathname === "/login") {
    if (sessionToken && refreshToken) {
      return NextResponse.redirect(new URL("/", request.nextUrl));
    } else {
      return NextResponse.next();
    }
  }
  if (!sessionToken || !refreshToken) {
    return redirectToLogin(request);
  }
}
