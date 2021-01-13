import React, { Component } from 'react';
import { Theme, createStyles, makeStyles, useTheme } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import Icon from '@material-ui/core/Icon';
import { css } from '@emotion/core';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import { Button, IconButton, Snackbar } from '@material-ui/core';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import TextField from '@material-ui/core/TextField';
import InputAdornment from '@material-ui/core/InputAdornment';
import SnackbarContent from '@material-ui/core/SnackbarContent';
import WarningIcon from '@material-ui/icons/Warning';
import clsx from 'clsx';
import { amber, green } from '@material-ui/core/colors';
import CheckCircleIcon from '@material-ui/icons/CheckCircle';
import ErrorIcon from '@material-ui/icons/Error';
import InfoIcon from '@material-ui/icons/Info';
import CloseIcon from '@material-ui/icons/Close';
import SaveIcon from '@material-ui/icons/Save';
import SearchIcon from '@material-ui/icons/Search';
import CircularProgress from '@material-ui/core/CircularProgress';
import Fab from '@material-ui/core/Fab';

//@ts-ignore
import Download from '@axetroy/react-download';
import { string } from 'prop-types';


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
      flex: '0.5',
      width: '100%'
    },
    wrapper: {
      margin: theme.spacing(1),
      position: 'relative',
    },
    cover: {
      backgroundSize: 'contain',
      width: 75,
      height: 75,
      margin: 'auto',
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
    },
    greenbutton: {
      backgroundColor: green[500],
      margin: theme.spacing(1),
    },
    fabProgress: {
      color: green[500],
      position: 'absolute',
      top: 2,
      left: 2,
      zIndex: 1,
    },
  }),
);
const useStyles1 = makeStyles(theme => ({
  success: {
    backgroundColor: green[600],
  },
  error: {
    backgroundColor: theme.palette.error.dark,
  },
  info: {
    backgroundColor: theme.palette.primary.main,
  },
  warning: {
    backgroundColor: amber[700],
  },
  icon: {
    fontSize: 20,
  },
  iconVariant: {
    opacity: 0.9,
    marginRight: theme.spacing(1),
  },
  message: {
    display: 'flex',
    alignItems: 'center',
  },
}));
const variantIcon = {
  success: CheckCircleIcon,
  warning: WarningIcon,
  error: ErrorIcon,
  info: InfoIcon,
};
const Proj = ({ data, onCreateVersion, onDownloadFile, files, loading, onShowDetails }: { data: any, onCreateVersion: any, onDownloadFile: any, files: Map<string, any>, loading: Map<string, boolean>, onShowDetails: any }) => {
  const override = css`
  float: 'right';
`;
  console.log(loading)
  const shouldComponentUpdate = () => {
    return false; // Will cause component to never re-render.
  }
  const classes = useStyles();
  console.log(data);
  if (!data) {
    return (
      <div> Loading ...</div>
    );
  }
  const mydata = data.projects;
  console.log(mydata)
  const iconlist = ["https://image.flaticon.com/icons/svg/86/86572.svg", "https://image.flaticon.com/icons/svg/86/86580.svg",
    "https://image.flaticon.com/icons/svg/813/813488.svg", "https://image.flaticon.com/icons/svg/813/813430.svg",
    "https://image.flaticon.com/icons/svg/813/813466.svg", "https://image.flaticon.com/icons/svg/813/813531.svg",
    "https://image.flaticon.com/icons/svg/2034/2034786.svg", "https://image.flaticon.com/icons/svg/813/813491.svg",
    "https://image.flaticon.com/icons/svg/81/81089.svg"]

  if (mydata) {
    console.log("toto")
    return (
      <div>
        <div className={classes.controls}>
          <Button
            variant="contained"
            color="primary"
            className={classes.button}
            onClick={() => onCreateVersion()}
            endIcon={<Icon>send</Icon>}>
            Create new Project
                  </Button>
        </div>

        <Grid container wrap="wrap" spacing={3} style={{ padding: '15px' }}>
          {Object.keys(mydata).map((d, key) => (
            <Grid item xs>
              <Card className={classes.card}>
                <div className={classes.details}>
                  <CardContent className={classes.content} >
                    <Typography component="h5" variant="h5">
                      {d}
                    </Typography>
                    <Table className={classes.table} size="small" aria-label="a dense table">
                      <TableHead>
                        <TableRow>
                          <TableCell>Versions</TableCell>
                          <TableCell>Generate</TableCell>
                          <TableCell>Details</TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {mydata[d].map((row: { name: any }) => (
                          <TableRow key={row.name}>
                            <TableCell component="th" scope="row">
                              {row.name}
                            </TableCell>
                            <TableCell>
                              {files.get(row.name) ?
                                <Download file={files.get(row.name).filename} content={files.get(row.name).contents}>
                                  <Fab size="small" color="primary" className={classes.greenbutton} aria-label="delete">
                                    <SaveIcon />
                                  </Fab>
                                </Download>
                                : <div className={classes.wrapper}>
                                  <Fab size="small" color="primary" onClick={() => onDownloadFile(row.name)} className={classes.button} aria-label="delete">
                                    <SaveIcon />
                                  </Fab>
                                  {loading.get(row.name) && <CircularProgress className={classes.fabProgress} size={52} />}
                                </div>
                              }

                            </TableCell>
                            <TableCell>
                              <Fab size="small" color="primary" onClick={() => onShowDetails(row.name)} className={classes.button} aria-label="delete">
                                <SearchIcon />
                              </Fab>
                            </TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  </CardContent>
                  <div className={classes.controls}>
                    <Button
                      variant="contained"
                      color="primary"
                      className={classes.button}
                      onClick={() => onCreateVersion(d)}
                      endIcon={<Icon>send</Icon>}>
                      Create version
                  </Button>
                  </div>
                </div>
                <CardMedia
                  className={classes.cover}
                  image={iconlist[(key % iconlist.length)]}
                  title={d}
                />
              </Card>
            </Grid>
          ))}
        </Grid>
      </div>
    );
  } else {
    return (
      <div> Loading ...</div>
    );
  }
}
function MySnackbarContentWrapper(props: any) {
  const classes = useStyles1();
  const { className, message, onClose, variant, ...other } = props;
  var Icon: any;
  var classe: any;
  if (variant === 'success') {
    Icon = variantIcon.success;
    classe = classes.success;
  } else if (variant === 'warning') {
    Icon = variantIcon.warning;
    classe = classes.warning;
  } else if (variant === 'error') {
    Icon = variantIcon.error;
    classe = classes.error;
  } else {
    Icon = variantIcon.info;
    classe = classes.info;
  }

  return (
    <SnackbarContent
      className={clsx(classe, className)}
      aria-describedby="client-snackbar"
      message={
        <span id="client-snackbar" className={classes.message}>
          <Icon className={clsx(classes.icon, classes.iconVariant)} />
          {message}
        </span>
      }
      action={[
        <IconButton key="close" aria-label="close" color="inherit" onClick={onClose}>
          <CloseIcon className={classes.icon} />
        </IconButton>,
      ]}
      {...other}
    />
  );
}

class Project extends Component<{}, { projects: any, showDialog: boolean, showSnack: boolean, version: string, project: string, files: Map<string, any>, loading: Map<string, boolean>, showDialogDetails: boolean, projectDetails: any }> {
  classes: any;
  theme: any;
  data: any;
  selectedChart: any;
  config: any;
  variant = ""
  message = ""
  createproject: string = ""
  state = {
    projects: [],
    showDialog: false,
    showSnack: false,
    version: "0",
    project: "",
    files: new Map<string, any>(),
    loading: new Map<string, boolean>(),
    showDialogDetails: false,
    projectDetails: []
  }


  render() {
    console.log("main dash rendering")
    return (
      <div>
        <Snackbar
          anchorOrigin={{
            vertical: 'top',
            horizontal: 'right',
          }}
          open={this.state.showSnack}
          autoHideDuration={6000}
          onClose={this.handleCloseSnack}
        >
          <MySnackbarContentWrapper
            onClose={this.handleCloseSnack}
            variant={this.variant}
            message={this.message}
          />
        </Snackbar>
        <Proj data={this.state} onCreateVersion={this.handleClickOpen} onDownloadFile={this.downloadFile} files={this.state.files} loading={this.state.loading} onShowDetails={this.onShowDetails} />
        <Dialog open={this.state.showDialog} onClose={this.handleCancel} aria-labelledby="form-dialog-title">
          <DialogTitle id="form-dialog-title">Creation of Project/Version {this.state.project}</DialogTitle>
          <DialogContent>
            <DialogContentText>
              {this.createproject ? "Please enter a version, it should not contain '_'" : "Please enter a project name , it should not contain '_'"}
            </DialogContentText>
            {this.createproject ?
              <TextField
                autoFocus
                margin="dense"
                id="name"
                label="Project name"
                value={this.state.version}
                type="text"
                onChange={this.handleChangeVersion}
                InputProps={{
                  startAdornment: <InputAdornment position="start">Project_{this.state.project}_Version_</InputAdornment>,
                }}
                fullWidth
              />
              : <TextField
                autoFocus
                margin="dense"
                id="name"
                label="Project name"
                type="text"
                value={this.state.project}
                onChange={this.handleChangeProject}
                InputProps={{
                  startAdornment: <InputAdornment position="start">Project_</InputAdornment>,
                }}
                fullWidth
              />}
          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleCancel} color="primary">
              Cancel
          </Button>
            <Button onClick={this.handleClose} color="primary">
              Create
          </Button>
          </DialogActions>
        </Dialog>
        <Dialog open={this.state.showDialogDetails} onClose={this.handleCancel} aria-labelledby="form-dialog-title">
          <DialogTitle id="form-dialog-title">List of versions</DialogTitle>
          <DialogContent>
            <DialogContentText>
              <Table  size="small" aria-label="a dense table">
                <TableHead>
                  <TableRow>
                    <TableCell>App</TableCell>
                    <TableCell>ChartVersion</TableCell>
                    <TableCell>Version</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {this.state.projectDetails.map((row: { metadata: any }) => (
                    <TableRow key={row.metadata.name}>
                      <TableCell component="th" scope="row">
                        {row.metadata.name}
                      </TableCell>
                      <TableCell>
                        {row.metadata.version}
                      </TableCell>
                      <TableCell>
                        {row.metadata.appVersion}
                      </TableCell>
                     
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </DialogContentText>
          </DialogContent>
          <DialogActions>
          <Button onClick={this.handleCancel} color="primary">
              Close
          </Button>
            <Button onClick={this.handleCancel} color="primary">
              Export to Excel
          </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }

  public onShowDetails = (project: string) => {
    fetch(window.location.origin+'/api/v1/getProjectDetail', {
   // fetch('http://172.24.34.179:8080/api/v1/getProjectDetail', {
      method: 'POST',
      body: JSON.stringify(project)
    }).then(res => res.json())
      .then((data) => {
        console.log(data)
        this.setState({ showDialogDetails: true, projectDetails: data.charts })
      })
  }

  public downloadFile = (project: string) => {
    var delToCreate = {
      label: project
    }
    this.state.loading.set(project, true)
    this.setState({ loading: this.state.loading })
    fetch(window.location.origin+'/api/v1/generateDelivery', {
   // fetch('http://172.24.34.179:8080/api/v1/generateDelivery', {
      method: 'POST',
      body: JSON.stringify(delToCreate)
    }).then(res => res.text())
      .then((data) => {
        console.log(data)
        this.message = 'OK'
        this.variant = 'success'
        this.setState({ showSnack: true })
        this.state.loading.set(project, false)
        this.state.files.set(project, {
          mime: 'text/html',
          filename: project + 'delivery-script-v1.0.sh',
          contents: data,
        })
      })
      .catch(console.log)
  }
  public handleChangeProject = (event: any) => {
    this.setState({ project: event.target.value });
  }
  public handleChangeVersion = (event: any) => {
    this.setState({ version: event.target.value });
  }
  public handleCloseSnack = (_event: any, reason: string) => {
    if (reason === 'clickaway') {
      return;
    }

    this.setState({ showSnack: false })
  };
  /*myChangeHandler = (event: any) => {
    this.version = event.target.value
  }*/
  public handleClickOpen = (event: any) => {
    console.log(event)
    this.createproject = event;
    this.setState({ project: event, showDialog: true })
  };

  public handleClose = (event: any) => {
    this.setState({ showDialog: false });
    var projectToCreate = "Project_" + this.state.project + "_Version_" + this.state.version
    console.log(projectToCreate)
    this.message = 'Sending'
    this.variant = 'info'
    this.setState({ showSnack: true })
    fetch(window.location.origin+'/api/v1/postProject', {
    //fetch('http://172.24.34.179:8080/api/v1/postProject', {
      method: 'POST',
      body: JSON.stringify(projectToCreate)
    }).then(
      response => {
        this.message = 'OK'
        this.variant = 'success'
        this.setState({ showSnack: true })
        this.loadData();
      },
      error => {
        this.message = error
        this.variant = 'error'
        this.setState({ showSnack: true })
      }
    )
  };
  public handleCancel = (event: any) => {
    this.setState({ showDialog: false, showDialogDetails: false });
  };

  public loadData = () => {
    fetch(window.location.origin+'/api/v1/getProjects')
    //fetch('http://172.24.34.179:8080/api/v1/getProjects')
      .then(res => res.json())
      .then((data) => {
        this.data = data;
        this.setState({ projects: data });
        console.log(data);
      })
      .catch(console.log)
  }

  componentDidMount() {
    console.log("fetch")
    // this.setState({charts: jsonData});
    // Get the config from the server too
    this.loadData();
  }
}
export default Project;
