import { Heading1, Text } from '@/components/ui/typography';
import { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Overview - Serentin Dashboard',
  description: 'Employee Overview',
};

export default function OverviewPage() {
  return (
    <>
      <Heading1 className=''>Overview</Heading1>
      <Text>Employee Overview</Text>
    </>
  );
}
