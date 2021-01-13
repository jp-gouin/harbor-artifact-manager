import React from 'react';
import { Theme, createStyles, makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import { Button, Icon, ListItemAvatar, Avatar } from '@material-ui/core';
import Paper from '@material-ui/core/Paper';
import Checkbox from '@material-ui/core/Checkbox';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Radio from '@material-ui/core/Radio';
import SvgIcon from '@material-ui/core/SvgIcon';
import { ReactComponent as ApprovedLogo } from './assets/approved.svg';
import { ReactComponent as UnApprovedLogo } from './assets/unapproved.svg';
import MaterialTable from 'material-table'
import clsx from 'clsx';
import CheckCircleIcon from '@material-ui/icons/CheckCircle';
import ErrorIcon from '@material-ui/icons/Error';
import InfoIcon from '@material-ui/icons/Info';
import CloseIcon from '@material-ui/icons/Close';
import { amber, green } from '@material-ui/core/colors';
import IconButton from '@material-ui/core/IconButton';
import Snackbar from '@material-ui/core/Snackbar';
import SnackbarContent from '@material-ui/core/SnackbarContent';
import WarningIcon from '@material-ui/icons/Warning';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Chip from '@material-ui/core/Chip';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
      padding: '2em',
    },
    avatar: {
      backgroundColor: 'transparent',
      width: 40,
      height: 40,

    },
    logo: {
      width: 40,
      height: 40,
      '$root.Mui-focusVisible &': {
        outline: '2px auto rgba(19,124,189,.6)',
        outlineOffset: 2,
      },
      'input:hover ~ &': {
        backgroundColor: '#ebf1f5',
      },
      'input:disabled ~ &': {
        boxShadow: 'none',
        background: 'rgba(206,217,224,.5)',
      },
    },
    list: {
      width: '100%',
      maxWidth: 360,
      backgroundColor: theme.palette.background.paper,
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
    button: {
      margin: theme.spacing(1),
    },
    leftgrid: {
      flexBasis: '0%',
    },
    chips: {
      display: 'flex',
      flexWrap: 'wrap',
    },
    chip: {
      margin: 2,
    },
    noLabel: {
      marginTop: theme.spacing(3),
    },
    listitem: {
      flexFlow: 'wrap',
    }
  }),
);
const variantIcon = {
  success: CheckCircleIcon,
  warning: WarningIcon,
  error: ErrorIcon,
  info: InfoIcon,
};

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

const MyChart = ({ value, config, onReturn }: { value: any, config: any, onReturn: any }) => {

  console.log(value)
  console.log(config)
  const [open, setOpen] = React.useState({ show: false, message: "", variant: "" });
  const handleClose = (_event: any, reason: string) => {
    if (reason === 'clickaway') {
      return;
    }

    setOpen({ show: false, message: "", variant: "" });
  };
  const ITEM_HEIGHT = 48;
  const ITEM_PADDING_TOP = 8;
  const MenuProps = {
    PaperProps: {
      style: {
        maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
        width: 250,
      },
    },
  };
  value.charts.sort(function (a: any, b: any) {
    return a.version - b.version;
  })
  value.allDockerImages.sort(function (a: any, b: any) {
    var left: any = a.repository + ":" + a.tag;
    var right: any = b.repository + ":" + b.tag;
    return left - right;
  })
  const isChartValid = (c: any) => {
    if (!c || c.length === 0)
      return false;
    var found = false;
    c.forEach((l: { name: string, deleted: boolean }) => {
      if (l.name === config.configlabel.name && !l.deleted) {
        found = true;
        return
      }
    });
    return found;
  }
  const send = (data:any, url = '/api/v1/postChartData') => {
    console.log(data);
    setOpen({ show: true, message: 'Sending validation', variant: 'info' })
    fetch(window.location.origin+url, {
    //  fetch("http://172.24.34.179:8080"+url, {
      method: 'POST',
      body: JSON.stringify(data)
    }).then(
      response => setOpen({ show: true, message: 'OK', variant: 'success' }),

      error => setOpen({ show: true, message: error, variant: 'error' })
    )

  }
  const [selectedValue, setSelectedValue] = React.useState('0');
  const classes = useStyles();
  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSelectedValue(event.target.value);
    console.log("selected")
    console.log(event.target.value)
  };

  const initChartValidated: number[] = [];
  const initCurDockValidated: string[] = [];
  const initAllDockerValidated: string[] = [];

  console.log('init state')
  for (var i = 0; i < value.allDockerImages.length; i++) {
    if (isChartValid(value.allDockerImages[i].labels)) {
      initAllDockerValidated.push(value.allDockerImages[i].repository + ' : ' + value.allDockerImages[i].tag);
    }
  }
  for (i = 0; i < value.charts.length; i++) {
    if (isChartValid(value.charts[i].labels)) {
      initChartValidated.push(i);
    }
    for (var j = 0; j < value.charts[i].currentDockerImages.length; j++) {
      if (isChartValid(value.charts[i].currentDockerImages[j].labels)) {
        initCurDockValidated.push(value.charts[i].currentDockerImages[j].repository + ' : ' + value.charts[i].currentDockerImages[j].tag)
      }
    }
  }

  const handleToogle = (v: any, list: any, method: any, obj: any, di: boolean) => () => {
    //list of label directly passed as obj
    console.log("toogle")
    console.log(obj)
    if (!obj.labels || obj.labels.length === 0) {
      console.log("init list")
      obj.labels = [];
      obj.labels.push(config.configlabel);
    } else {
      var index = -1;
      for (var i = 0; i < obj.labels.length; i++) {
        console.log(obj.labels[i]);
        if (obj.labels[i].name === config.configlabel.name) {
          console.log(i);
          index = i;
          obj.labels[i].deleted = obj.labels[i].deleted ? false : true;
        }
      }
      /*
      if (index === -1) {
        console.log("push")
        obj.labels.push(config.configlabel);
      } else {
        // obj.labels[i]
      }*/
    }

    if(di){
     let data = {
        name: value.name,
        allDockerImages: [obj],
        charts: [],
        project: value.project
      }
      send(data)
    }else {
      let data = {
        name: value.name,
        allDockerImages: [],
        charts: [obj],
        project: value.project
      }
      send(data)
    }

    console.log(value);
    const currentIndex = list.indexOf(v);
    const newChecked = [...list];

    if (currentIndex === -1) {
      newChecked.push(v);
    } else {
      newChecked.splice(currentIndex, 1);
    }
    method(newChecked);
  }

  const [chartValidated, setchartValidated] = React.useState(initChartValidated);
  // const [curDockValidated, setcurDockValidated] = React.useState(initCurDockValidated);
  const [allDockValidated, setAllDockValidated] = React.useState(initAllDockerValidated);

  const [personName, setPersonName] = React.useState<[]>([]);

  const handleSelectChange = (event: any, di: any, docki: boolean) => {
    console.log(event.target.value)
    let projectName = "";
    let toDelete = false;
    event.target.value.forEach((element2: any) => {
      var found = false;
      di.labels.forEach((element: any) => {
        if (element.name === element2.name) {
          found = true;
          console.log("found")
          element.deleted = false;
          projectName = element2.name;
        }
      });
      if (!found) {
        console.log("push")
        di.labels.push({
          id: element2.id,
          creation_time: element2.creation_time,
          deleted: element2.deleted,
          description: element2.description,
          name: element2.name,
          project_id: element2.project_id,
          scope: element2.scope,
          update_time: element2.update_time
        })
        projectName = element2.name;
      }
    });
    di.labels.forEach((element: any) => {
      var found = false;
      console.log(element)
      event.target.value.forEach((element2: any) => {
        console.log(element2)
        if (element.name === element2.name) {
          found = true;
        }
      });
      if (!found && element.name !== "ADS_Validated") {
        console.log("not found")
        element.deleted = true;
        toDelete = true;
        projectName = element.name;
      }
    });
    console.log(di);
    let data;
    if(docki){
      data = {
         name: value.name,
         allDockerImages: [di],
         charts: [],
         project: value.project,
         projectLab: projectName
       }
     }else {
       data = {
         name: value.name,
         allDockerImages: [],
         charts: [di],
         project: value.project,
         projectLab: projectName
       }
     }
     send(data, toDelete ? '/api/v1/removeProjectToArtifact' : '/api/v1/addProjectToArtifact')

    /*if (isChartValid(di.labels)) {
      di.labels = []
      di.labels.push(config.configlabel);
      di.labels.push.apply(di.labels, event.target.value)

    } else {
      di.labels = []
      di.labels.push.apply(di.labels, event.target.value)
    }*/
    setPersonName(event.target.value as []);
  };
  const getProject = (d: any) => {
    var r: any[] = [];
    if (!d.labels) {
      d.labels = []
    }
    config.projects.forEach((element: any) => {
      if (containVersion(element, d.labels)) {
        r.push(element);
      }
    });
    return r;
  }
  const containVersion = (e: any, d: any) => {
    var found = false;
    d.forEach((element: any) => {
      if (e.name === element.name && ! element.deleted) {
        found = true
        return true;
      }
    });
    return found;
  }

  /* const handleChangeMultiple = (event: React.ChangeEvent<{ value: unknown }>) => {
     const { options } = event.target as HTMLSelectElement;
     const value = [];
     for (let i = 0, l = options.length; i < l; i += 1) {
       if (options[i].selected) {
         value.push(options[i].value);
       }
     }
     setPersonName(value);
   };*/
  if (!value) {
    return (
      <div> Loading ...</div>
    );
  }

  return (
    <div className={classes.root}>
      <Snackbar
        anchorOrigin={{
          vertical: 'top',
          horizontal: 'right',
        }}
        open={open.show}
        autoHideDuration={6000}
        onClose={handleClose}
      >
        <MySnackbarContentWrapper
          onClose={handleClose}
          variant={open.variant}
          message={open.message}
        />
      </Snackbar>
      <Button
        variant="contained"
        color="primary"
        className={classes.button}
        onClick={onReturn}
        startIcon={<Icon>send</Icon>}>
        Return
    </Button>
      {/*<Button
        variant="contained"
        color="primary"
        className={classes.button}
        onClick={send}
        endIcon={<Icon>check</Icon>}>
        Validate
      </Button>:*/}
      <Typography variant="h2" component="h2">
        {value.name}
      </Typography>
      <Grid container justify="center" direction="row" spacing={3}>
        <Grid
          item
          spacing={3}
          container
          justify="flex-start"
          direction="column"
          xs={4}
        >
          <Grid item xs={12} className={classes.leftgrid}>
            <Paper className={classes.paper}>
              <Typography component="h5" variant="h5">
                Charts
                  </Typography>
              <List className={classes.list} >
                {Object.keys(value.charts).map((d, key) => (
                  <ListItem className={classes.listitem}>
                    <ListItemAvatar>
                      <Avatar className={classes.avatar}>
                        <Checkbox
                          edge="end"
                          onChange={handleToogle(key, chartValidated, setchartValidated, value.charts[d], false)}
                          checked={chartValidated.indexOf(key) !== -1}
                          inputProps={{ 'aria-label': 'decorative checkbox' }}
                          checkedIcon={<ApprovedLogo className={classes.logo} />}
                          icon={<UnApprovedLogo className={classes.logo} />}
                        />
                      </Avatar>
                    </ListItemAvatar>
                    <ListItemText primary={"Chart version : " + value.charts[d].version} secondary={"App version : " + value.charts[d].metadata.appVersion} />
                    <Radio
                      checked={selectedValue === d}
                      onChange={handleChange}
                      value={d}
                      name="radio-button-demo"
                      inputProps={{ 'aria-label': d }}
                    />
                    <Select
                      labelId="demo-mutiple-chip-label"
                      id="demo-mutiple-chip"
                      multiple
                      value={getProject(value.charts[d])}
                      onChange={(e) => handleSelectChange(e, value.charts[d], false)}
                      input={<Input id="select-multiple-chip" />}
                      renderValue={(selected: any) => (
                        <div className={classes.chips}>
                          {Object.keys(selected).map((v, key) => (
                            <Chip key={key} label={selected[v].name} className={classes.chip}  color="primary"/>
                          ))}
                        </div>
                      )}
                      MenuProps={MenuProps}
                    >
                      {config.projects.map((label: any, key: any) => {
                        return (<MenuItem key={key} value={label}>
                          {label.name}
                        </MenuItem>);
                      })}
                    </Select>
                  </ListItem>
                ))}
              </List>
            </Paper>
          </Grid>
          <Grid item xs={12} className={classes.leftgrid}>
            <Paper className={classes.paper}>
              <Typography component="h5" variant="h5">
                Docker images directly used by the chart
                  </Typography>
              <List dense className={classes.list}>
                {Object.keys(value.charts[selectedValue].currentDockerImages).map((c, key) => {
                  return (
                    <ListItem >
                      <ListItemAvatar>
                        <Checkbox
                          edge="end"
                          disabled
                          onChange={handleToogle(value.charts[selectedValue].currentDockerImages[c].repository + ' : ' + value.charts[selectedValue].currentDockerImages[c].tag, allDockValidated, setAllDockValidated, value.charts[selectedValue].currentDockerImages[c],true)}
                          checked={allDockValidated.indexOf(value.charts[selectedValue].currentDockerImages[c].repository + ' : ' + value.charts[selectedValue].currentDockerImages[c].tag) !== -1}
                          inputProps={{ 'aria-label': 'decorative checkbox' }}
                          checkedIcon={<ApprovedLogo className={classes.logo} />}
                          icon={<UnApprovedLogo className={classes.logo} />}
                        />
                      </ListItemAvatar>
                      <ListItemText primary={value.charts[selectedValue].currentDockerImages[c].repository + ' : ' + value.charts[selectedValue].currentDockerImages[c].tag} />
                    </ListItem>
                  );
                })}
              </List>
            </Paper>
          </Grid>
          <Grid item xs={12} className={classes.leftgrid}>
            <Paper className={classes.paper}>
              <Typography component="h5" variant="h5">
                Dependencies
                  </Typography>
              <List dense className={classes.list}>
                {Object.keys(value.charts[selectedValue].dependencies).map((c, key) => {
                  return (
                    <ListItem >
                      <Chip
        label={value.charts[selectedValue].dependencies[c].name + ' : ' + value.charts[selectedValue].dependencies[c].version}
        clickable
        color="primary"
      />
                    </ListItem>
                  );
                })}
              </List>
            </Paper>
          </Grid>
        </Grid>
        <Grid item xs={8}>
          <Paper className={classes.paper}>
            <div style={{ maxWidth: '100%' }}>
              <MaterialTable
                options={{
                  pageSize: 20
                }}
                columns={[
                  {
                    title: 'Validated', field: 'validated', cellStyle: { padding: '5px' }, render: rowData =>
                      <Checkbox
                        edge="end"
                        onChange={handleToogle(rowData.repository + ' : ' + rowData.tag, allDockValidated, setAllDockValidated, rowData,true)}
                        checked={allDockValidated.indexOf(rowData.repository + ' : ' + rowData.tag) !== -1}
                        inputProps={{ 'aria-label': 'decorative checkbox' }}
                        checkedIcon={<ApprovedLogo className={classes.logo} />}
                        icon={<UnApprovedLogo className={classes.logo} />}
                      />
                  },
                  { title: 'Repo', field: 'repository', cellStyle: { padding: '5px' } },
                  { title: 'Tag', field: 'tag', cellStyle: { padding: '5px' } },
                  { title: 'Severity', field: 'scan_overview.severity', type: 'numeric', cellStyle: { padding: '5px' } },
                  {
                    title: 'Version', field: 'projects', render: rowData =>
                      <Select
                        labelId="demo-mutiple-chip-label"
                        id="demo-mutiple-chip"
                        multiple
                        value={getProject(rowData)}
                        onChange={(e) => handleSelectChange(e, rowData, true)}
                        input={<Input id="select-multiple-chip" />}
                        renderValue={(selected: any) => (
                          <div className={classes.chips}>
                            {Object.keys(selected).map((v, key) => (
                              <Chip key={key} label={selected[v].name} className={classes.chip} color="primary" />
                            ))}
                          </div>
                        )}
                        MenuProps={MenuProps}
                      >
                        {config.projects.map((label: any, key: any) => {
                          return (<MenuItem key={key} value={label}>
                            {label.name}
                          </MenuItem>);
                        })}
                      </Select>
                  }

                ]}
                data={value.allDockerImages}
                title="All docker images"
              />
            </div>
          </Paper>
        </Grid>
      </Grid>
    </div >
  );
}

export default MyChart;
