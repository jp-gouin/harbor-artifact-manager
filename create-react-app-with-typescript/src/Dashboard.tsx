import React, { Component } from 'react';
import { Theme, createStyles, makeStyles, useTheme } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import Icon from '@material-ui/core/Icon';
import ReplayIcon from '@material-ui/icons/Replay';
import { css } from '@emotion/core';
import RingLoader from 'react-spinners/RingLoader';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import MyChart from './MyChart';
import { Button, IconButton } from '@material-ui/core';
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    card: {
      display: 'flex',
      height: '100%',
    },
    details: {
      display: 'flex',
      flexDirection: 'column',
    },
    content: {
      flex: '1',
      width: '100%'
    },
    cover: {
      width: 75,
      height: 75,
      backgroundSize: 'contain'
    },
    controls: {
      display: 'flex',
      alignItems: 'center',
      paddingLeft: theme.spacing(1),
      paddingBottom: theme.spacing(1),
    },
    playIcon: {
      height: 38,
      width: 38,
    },
    button: {
      margin: theme.spacing(1),
    },
    table: {
      maxWidth: 150,
    },
    loader: {
      float: 'right',
    }
  }),
);


const Dash = ({ data, onSelectChart, getChartData }: { data: any, onSelectChart: any, getChartData: any }) => {
  const override = css`
  float: 'right';
`;

  const classes = useStyles();
  console.log(data);
  if (!data) {
    return (
      <div> Loading ...</div>
    );
  }
  const mydata = data.charts;
  console.log(mydata)
  if (mydata) {
    console.log("toto")
    return (
      <Grid container spacing={3} style={{ padding: '15px' }}>
        {Object.keys(mydata).map((d, key) => (
          <Grid item xs={3}>

            <Card className={classes.card}>
              <div className={classes.details}>
                <CardContent className={classes.content} >
                  <Grid container justify="space-around" alignItems="center" spacing={3}>
                    <Grid item xs={4}>
                      <CardMedia
                        className={classes.cover}
                        image={mydata[d].icon}
                        title={mydata[d].name}
                      />
                    </Grid>
                    <Grid item xs={4}>
                      <Typography component="h5" variant="h5">
                        {mydata[d].name}
                      </Typography>
                    </Grid>
                    <Grid item xs={4}>
                      <RingLoader 
                        css={override}
                        sizeUnit={"px"}
                        size={30}
                        color={'#123abc'}
                        loading={mydata[d].charts ? false : true}
                      />
                    </Grid>
                  </Grid>
                  <Typography variant="subtitle1" color="textSecondary">
                    Latest version : {mydata[d].latest_version}
                  </Typography>
                  {mydata[d].charts ? (
                    <Table className={classes.table} size="small" aria-label="a dense table">
                      <TableHead>
                        <TableRow>
                          <TableCell>ChartVersion</TableCell>
                          <TableCell align="left">AppVersion</TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {mydata[d].charts.map((row: { version: any; metadata: any; }) => (
                          <TableRow key={row.version}>
                            <TableCell component="th" scope="row">
                              {row.version}
                            </TableCell>
                            <TableCell align="left">{row.metadata.appVersion}</TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  ) : (
                      <Typography variant="subtitle1" color="textSecondary">
                        Total Charts : {mydata[d].total_versions}
                      </Typography>
                    )}
                  {mydata[d].charts ? (
                    <Typography variant="subtitle1" color="textSecondary">
                      {mydata[d].allDockerImages?mydata[d].allDockerImages.length:0} Docker images
                </Typography>) : ''}
                </CardContent>
                <div className={classes.controls}>
                  <Button
                    variant="contained"
                    disabled={mydata[d].charts ? false: true}
                    color="primary"
                    className={classes.button}
                    onClick={() => onSelectChart(mydata[d])}
                    endIcon={<Icon>send</Icon>}>
                    Validate
                  </Button>
                  <IconButton className={classes.button} onClick={() => getChartData(mydata[d])} aria-label="delete">
                    <ReplayIcon />
                  </IconButton>
                </div>
              </div>

            </Card>
          </Grid>
        ))}
      </Grid>
    );
  } else {
    return (
      <div> Loading ...</div>
    );
  }


}

class Dashboard extends Component<{}, { charts: any, showResults: boolean }> {
  classes: any;
  theme: any;
  myData = [];
  selectedChart: any;
  config: any;

  public handleChartSelection = (event: any) => {
    console.log(event);
    this.selectedChart = event;
    this.setState({ showResults: true });
  }
  public handleReturn = (event: any) => {
    console.log(event);
    this.selectedChart = null;
    this.setState({ showResults: false });
  }
  render() {
    console.log("main dash rendering")
    if (!this.selectedChart) {
      return (
        <Dash data={this.state} onSelectChart={this.handleChartSelection} getChartData={this.getChartData} />
      );
    } else {
      return (
        <MyChart value={this.selectedChart} config={this.config} onReturn={this.handleReturn} />
      )
    }
  }
  public getChartData(d: any) {
    console.log(JSON.stringify(d));

    fetch(window.location.origin+'/api/v1/getChartData', {
      method: 'POST',
      body: JSON.stringify(d)
    })
      .then(res => res.json())
      .then((data) => {
        this.setState((state) => {
          state.charts.forEach((c: any, i: any) => {
            if (c.name === data.name) {
              console.log(c);
              data.charts.sort(function (a: any, b: any) {
                return a.version - b.version;
              })
              state.charts[i] = data
            }
          });

          return { charts: state.charts }
        });
      })
      .catch(console.log)
  }
  componentDidMount() {
    console.log("fetch")
    // this.setState({charts: jsonData});
    // Get the config from the server too
    fetch(window.location.origin+'/api/v1/getConfig')
    //fetch('http://172.24.34.179:8080/api/v1/getConfig')
      .then(res => res.json())
      .then((data) => {
        this.config = data
        console.log(data);
      })
      .catch(console.log)
    fetch(window.location.origin+'/api/v1/getChartList')
    //fetch('http://172.24.34.179:8080/api/v1/getChartList')
      .then(res => res.json())
      .then((data) => {
        data.sort(function (a: any, b: any) {
          return a.name.localeCompare(b.name);
        })
        this.setState({ charts: data });
        console.log(data);

      })
      .catch(console.log)
  }
}
export default Dashboard;
