// src/components/PRsOpenLongerThan5Days.js
import React, { useEffect, useState } from 'react';
import { Card, CardContent, Typography, List, ListItem, Box, CircularProgress } from '@mui/material';
import axios from 'axios';

export default function PRsOpenLongerThan5Days() {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // Fetch the dummy data from the backend
    axios.get('http://localhost:8080/metrics/prs-open-longer-than-5-days')
      .then((res) => {
        setData(res.data);
        setLoading(false);
      })
      .catch((err) => {
        console.error('Error fetching PRs:', err);
        setLoading(false);
      });
  }, []);

  if (loading) {
    return (
      <Card>
        <CardContent>
          <Typography variant="h6">PRs Open Longer Than 5 Days</Typography>
          <Box sx={{ display: 'flex', justifyContent: 'center', padding: 2 }}>
            <CircularProgress />
          </Box>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card sx={{ height: '100%' }}>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          PRs Open Longer Than 5 Days
        </Typography>
        <List>
          {data.map((pr, idx) => (
            <ListItem key={idx}>
              <Box sx={{ display: 'flex', flexDirection: 'column' }}>
                <Typography variant="body1"><strong>Repo:</strong> {pr.repo_name}</Typography>
                <Typography variant="body2"><strong>PR #{pr.pr_number}</strong> - {pr.pr_title}</Typography>
                <Typography variant="body2"><strong>Opened by:</strong> {pr.developer}</Typography>
                <Typography variant="body2"><strong>Days Open:</strong> {pr.days_open}</Typography>
              </Box>
            </ListItem>
          ))}
        </List>
      </CardContent>
    </Card>
  );
}
