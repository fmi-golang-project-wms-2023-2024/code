import { NextRequest, NextResponse } from "next/server";

export async function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl;
  const url = request.nextUrl.clone();

  const token = request.cookies.get("accessToken");
  const publicRoutes = ["/login", "/register"];

  if (request.nextUrl.pathname.startsWith("/_next")) {
    return NextResponse.next();
  }

  if (!token && !publicRoutes.includes(pathname)) {
    url.pathname = "/login";
    return NextResponse.redirect(url);
  }

  if (token && pathname == "/") {
    url.pathname = "/dashboard";
    return NextResponse.redirect(url);
  }

  if (token && pathname == "/login") {
    url.pathname = "/dashboard";
    return NextResponse.redirect(url);
  }

  if (token && pathname == "/register") {
    url.pathname = "/dashboard";
    return NextResponse.redirect(url);
  }

  return NextResponse.next();
}
