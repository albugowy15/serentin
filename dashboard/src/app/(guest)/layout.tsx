import { ThemeProvider } from '@/components/theme-provider';
import '@/app/globals.css';
import { Inter } from 'next/font/google';
import { Large, Small } from '@/components/ui/typography';
import HomeNavigation from '@/app/(guest)/components/home-navigation';
import { ThemeToggle } from '@/components/theme-toggle';
import { Button } from '@/components/ui/button';
import Link from 'next/link';

const inter = Inter({ subsets: ['latin'] });

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang='en' suppressHydrationWarning>
      <body className={inter.className}>
        <ThemeProvider
          attribute='class'
          defaultTheme='system'
          enableSystem
          disableTransitionOnChange
        >
          <header className='flex justify-between px-8 items-center py-2 border-b'>
            <Large>
              <Link href='/'>Serentin</Link>
            </Large>
            <HomeNavigation />
            <div className='flex items-center gap-2'>
              <ThemeToggle />
              <Button variant='default'>Login</Button>
            </div>
          </header>
          {children}
          <footer className='px-8 py-4 border-t'>
            <Small className='text-foreground/70'>
              Built by{' '}
              <a
                className='underline underline-offset-4'
                href='https://github.com/albugowy15'
              >
                albugowy15
              </a>
              . The source code is available on{' '}
              <a
                className='underline underline-offset-4'
                href='https://github.com/albugowy15/serentin/dashboard'
              >
                Github
              </a>
            </Small>
          </footer>
        </ThemeProvider>
      </body>
    </html>
  );
}
