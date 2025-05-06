// src/components/MainDashboard.js
import React from 'react';
import { Grid } from '@mui/material';
import TopCommitters from './TopCommitters';
// import TopContributors from './TopContributors';
// import TopReviewers from './TopReviewers';

export default function MainDashboard() {
  return (
    <Grid container spacing={2} sx={{ padding: 2 }}>
      <Grid item xs={12} md={6}>
        <TopCommitters />
      </Grid>
        {/* <Grid item xs={12} md={6}>
          <TopContributors />
        </Grid>
        <Grid item xs={12} md={6}>
          <TopReviewers />
        </Grid> */}
    </Grid>
  );
}
