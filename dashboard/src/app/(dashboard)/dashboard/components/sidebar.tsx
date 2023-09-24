'use client';

import { sidebarMenus } from '@/config/sidebar';
import { usePathname } from 'next/navigation';
import Link from 'next/link';
import { cn } from '@/lib/utils';

export default function Sidebar() {
  const pathname = usePathname();
  console.log(pathname);
  return (
    <aside className='px-4 pt-14'>
      <nav>
        <ul className='flex flex-col gap-3 mx-3 my-4'>
          {sidebarMenus.map((menu, index) => (
            <li key={index}>
              <Link
                href={menu.url}
                aria-label={menu.title}
                aria-labelledby={menu.title}
              >
                <div
                  className={cn(
                    'flex gap-5 items-center px-4 py-3 rounded-lg',
                    [
                      pathname === menu.url &&
                        'bg-primary text-primary-foreground',
                    ],
                  )}
                >
                  <menu.icon size={20} />
                  <span className='font-bold'>{menu.title}</span>
                </div>
              </Link>
            </li>
          ))}
        </ul>
      </nav>
    </aside>
  );
}
