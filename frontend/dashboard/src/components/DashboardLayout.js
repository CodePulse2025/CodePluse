import React from 'react';
import { Box } from '@mui/material';
import Header from './Header';
import Sidebar from './Sidebar';
import RepoGraph from './RepoGraph';
import TopReviewers from './TopReviewers';
import Contributors from './Contributors';
import PRsOpenLongerThan5Days from './PRsOpenLongerThan5Days';

export default function DashboardLayout() {
  return (
    <Box sx={{ display: 'flex' }}>
      <Sidebar />

      <Box sx={{ flexGrow: 1 }}>
        <Header />

        <Box sx={{ p: 3 }}>
          <RepoGraph />
          <Box sx={{ display: 'flex', gap: 2, mt: 2 }}>
            <TopReviewers />
            <Contributors />
            <PRsOpenLongerThan5Days />
          </Box>
          <Box sx={{ mt: 2 }}>
            
          </Box>
        </Box>
      </Box>
    </Box>
  );
}
