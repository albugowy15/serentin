import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import Sidebar from '@/app/(dashboard)/dashboard/components/sidebar';
import { ThemeProvider } from '@/components/theme-provider';
import { Inter } from 'next/font/google';
import '@/app/globals.css';

const inter = Inter({ subsets: ['latin'] });

export default function DashboardLayout({
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
          <div className='flex h-screen'>
            <Sidebar />
            <main className='bg-secondary flex-1 relative'>
              <section className='absolute right-0 top-0 p-4'>
                <Avatar>
                  <AvatarImage src='https://github.com/albugowy15.png' />
                  <AvatarFallback>Owi</AvatarFallback>
                </Avatar>
              </section>
              <div className='px-4 pt-16'>{children}</div>
            </main>
          </div>
        </ThemeProvider>
      </body>
    </html>
  );
}
