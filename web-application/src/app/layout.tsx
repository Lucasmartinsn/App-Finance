"use client"
import { Inter } from "next/font/google";
import GlobalStyle from "./globals";
import { ThemeProvider } from "styled-components";
import { darkTheme } from "@/styles/themes/dark";

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <head>
        <title>Application of Fiance</title>
        <link rel="shortcut icon" href="favicon.ico" type="image/x-icon" />
      </head>
      <ThemeProvider theme={darkTheme}>
        <GlobalStyle />
        <body className={inter.className}>{children}</body>
      </ThemeProvider>
    </html>
  );
}
