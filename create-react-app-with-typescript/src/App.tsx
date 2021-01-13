import React from 'react';
import Container from '@material-ui/core/Container';
import AppBar from './AppBar'

export default function App() {
  return (
    <Container maxWidth={false} style={{ padding: '0px 0px 0px 0px' }}>
        <AppBar />
    </Container>
  );
}
