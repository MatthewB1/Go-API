import React, { Component } from 'react';

import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';

import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import ButtonBase from '@material-ui/core/ButtonBase';

import Breadcrumbs from '@material-ui/lab/Breadcrumbs';
import Icon from '@material-ui/core/Icon'

import Dialog from '@material-ui/core/Dialog';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';


import Avatar from '@material-ui/core/Avatar';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import ListItemText from '@material-ui/core/ListItemText';
import PersonIcon from '@material-ui/icons/Person';

import decode from 'jwt-decode';


import ProjectComponent from './project'

const styles = theme => ({
    root: {
        flexGrow: 1,
    },
    breadcrumbs: {
        justifyContent: 'center',
        flexWrap: 'wrap',
    },
    crumb: {
        padding: `${theme.spacing.unit}px ${theme.spacing.unit * 2}px`,
    },
    paper: {
        height: 180,
        width: 200,
    },
    control: {
        padding: theme.spacing.unit * 2,
    },
    btn: {
        margin: 'auto',
        display: 'block',
        maxWidth: '100%',
        maxHeight: '100%',
    },
    icon: {
        margin: theme.spacing.unit * 2,
    },
});

class DashboardComponent extends Component {
    constructor(props) {
        super(props);
        this.state = {
            spacing: '16',
            username: '',
            accessLevel: '',

            error: null,
            newProject: false,
            selectedProject: null,

            users: [],
            projects: [],

            formError: null,
            projectname: '',
            projectlead: null
        };
    }

    componentDidMount() {
        const token = decode(localStorage.getItem('token'));
        this.setState({ username: token.name, accessLevel: token.admin ? ("admin") : ("user") })

        this.getData(token.name);
    }

    getData(username) {
        fetch('/api/projectAdministration/usersProjects?user=' + username, { method: 'GET' })
            .then(data => data.json())
            .then(res => {
                if (res.Success) {
                    if (res.Data == null)
                        this.setState({ projects: [] })
                    else
                        this.setState({ projects: res.Data });
                }
                else {
                    this.setState({ error: res.Error });
                }
            }
            );

        fetch('/api/userAdministration/users', { method: 'GET' })
            .then(data => data.json())
            .then(res => {
                if (res.Success) {
                    if (res.Data == null)
                        this.setState({ users: [] })
                    else
                        this.setState({ users: res.Data });
                }
                else {
                    this.setState({ error: res.Error });
                }
            }
            );
    }

    selectProject(value) {
        this.setState({ selectedProject: value })
    }

    handleChange = name => event => {
        this.setState({ [name]: event.target.value });
    };

    changeProjectLead(user) {
        this.setState({ projectlead: user });
    }



    newProject() {
        this.setState({ newProject: true })
    }

    handleClose = () => {
        this.setState({ projectname: '', projectlead: {}, newProject: false });
    };

    submitProject = () => {
        console.dir(this.state)
        //fetch new project
        if (this.state.projectname != '' &&  this.state.projectlead != null){
        fetch('/api/projectAdministration/project', { method: 'POST', body: JSON.stringify({ projectname: this.state.projectname, projectlead: this.state.projectlead }) })
            .then(data => data.json())
            .then(res => {
                if (!res.Success) {
                    alert(res.Error)
                }
                else {
                    //reload component
                    this.getData(this.state.username);
                    this.setState({ projectname: '', projectlead: null, newProject: false });
                }
            }
            );

        } else {
            if(this.state.projectlead == null)
                alert("please select a project lead!")
            if(this.state.projectname == '')
                alert("please provide a project name!")
        }
    };

    noProjects() {
        console.log(this.state.projects.length)
        return (this.state.projects.length === 0);
    }


    render() {
        const { classes } = this.props;
        const { spacing } = this.state;

        switch (this.state.accessLevel) {
            case "user": {
                if (this.state.error === null) {
                    if (this.state.selectedProject === null) {
                        if (this.state.projects.length !== 0) {
                            return (
                                <div>
                                    <div className={classes.breadcrumbs}>
                                        <Paper className={classes.crumb}>
                                            <Breadcrumbs arial-label="Breadcrumb">
                                                <Typography color="inherit">dashboard</Typography>
                                            </Breadcrumbs>
                                        </Paper>
                                        <br />
                                    </div>
                                    <Grid container className={classes.root} spacing={16}>
                                        <div style={{ "marginLeft": "auto", "marginRight": "auto" }}>
                                            <h2>Hi {this.state.username}! Your projects: </h2>
                                        </div>
                                        <Grid item xs={12}>
                                            <Grid container justify="center" spacing={Number(spacing)}>
                                                {this.state.projects.map(value => (
                                                    <Grid key={value.Projectname} item>
                                                        <ButtonBase onClick={() => this.selectProject(value)} className={classes.btn}>
                                                            <Paper className={classes.paper}>
                                                                <Typography gutterBottom variant="display2">
                                                                    {value.Projectname}
                                                                </Typography>
                                                            </Paper>
                                                        </ButtonBase>
                                                    </Grid>
                                                ))}
                                            </Grid>
                                        </Grid>
                                    </Grid>
                                </div>
                            );
                        } else {
                            return (
                                <div>
                                    <div className={classes.breadcrumbs}>
                                        <Paper className={classes.crumb}>
                                            <Breadcrumbs arial-label="Breadcrumb">
                                                <Typography color="inherit">dashboard</Typography>
                                            </Breadcrumbs>
                                        </Paper>
                                        <br />
                                    </div>
                                    <Grid container className={classes.root} spacing={16}>
                                        <div style={{ "marginLeft": "auto", "marginRight": "auto" }}>
                                            <h2>Hi {this.state.username}! Your projects: </h2>
                                        </div>
                                        <Grid item xs={12}>
                                            <Grid container justify="center" spacing={Number(spacing)}>
                                                <Grid item>
                                                    <Typography gutterBottom variant="display2">
                                                        no projects :~(
                                            </Typography>
                                                </Grid>
                                            </Grid>
                                        </Grid>
                                    </Grid>
                                </div>
                            );
                        }
                    }
                    else {
                        return (
                            <ProjectComponent project={this.state.selectedProject} />
                        )
                    }
                }
                else {
                    return (
                        <Grid container className={classes.root} spacing={16}>
                            <div style={{ "margin-left": "auto", "margin-right": "auto" }}>
                                <h2>Error loading your projects....</h2>
                                <p>{this.state.error}</p>
                            </div>
                        </Grid>
                    )
                }
            } //case user
            case "admin": {
                console.log("loading admin view!")
                if (this.state.error === null) {
                    if (this.state.selectedProject === null) {
                        return (
                            <div>
                                <div className={classes.breadcrumbs}>
                                    <Paper className={classes.crumb}>
                                        <Breadcrumbs arial-label="Breadcrumb">
                                            <Typography color="inherit">dashboard</Typography>
                                        </Breadcrumbs>
                                    </Paper>
                                    <br />
                                </div>
                                <Grid container className={classes.root} spacing={16}>
                                    <div style={{ "marginLeft": "auto", "marginRight": "auto" }}>
                                        <h2>Hi {this.state.username}! Your projects: </h2>
                                    </div>
                                    <Grid item xs={12}>
                                        <Grid container justify="center" spacing={Number(spacing)}>
                                            <Grid key="newProject" item>
                                                <ButtonBase onClick={() => this.newProject()} className={classes.btn}>
                                                    <Paper className={classes.paper}>
                                                        <Typography gutterBottom variant="display1">
                                                            New project
                                                        </Typography>
                                                        <div>
                                                            <Icon className={classes.icon} color="disabled" fontSize="large">
                                                                add_circle
                                                            </Icon>
                                                        </div>
                                                    </Paper>
                                                </ButtonBase>
                                                <Dialog
                                                    open={this.state.newProject}
                                                    onClose={this.handleClose}
                                                    aria-labelledby="form-dialog-title"
                                                >
                                                    <DialogTitle id="form-dialog-title">New Project</DialogTitle>
                                                    <DialogContent>
                                                        <DialogContentText>
                                                            Project name :
                                                        </DialogContentText>
                                                        <div>
                                                            <TextField
                                                                autoFocus
                                                                margin="dense"
                                                                id="projectname"
                                                                type="text"
                                                                fullwidth="true"
                                                                onChange={this.handleChange('projectname')}
                                                            />
                                                            <DialogContentText>
                                                                Project lead :
                                                            </DialogContentText>
                                                            <List>
                                                                {this.state.users.map(value => (
                                                                    <ListItem button onClick={() => this.changeProjectLead(value)} key={value.username}>
                                                                        <ListItemAvatar>
                                                                            <Avatar className={classes.avatar}>
                                                                                <PersonIcon />
                                                                            </Avatar>
                                                                        </ListItemAvatar>
                                                                        <ListItemText primary={value.username} />
                                                                    </ListItem>
                                                                ))}
                                                            </List>
                                                        </div>
                                                    </DialogContent>
                                                    <DialogActions>
                                                        <Button onClick={this.handleClose} color="primary">
                                                            Cancel
                                                    </Button>
                                                        <Button onClick={this.submitProject} color="primary">
                                                            Create
                                                    </Button>
                                                    </DialogActions>
                                                </Dialog>
                                            </Grid>
                                            {this.state.projects.map(value => (
                                                <Grid key={value} item>
                                                    <ButtonBase onClick={() => this.selectProject(value)} className={classes.btn}>
                                                        <Paper className={classes.paper}>
                                                            <Typography gutterBottom variant="display2">
                                                                {value.Projectname}
                                                            </Typography>
                                                        </Paper>
                                                    </ButtonBase>
                                                </Grid>
                                            ))}
                                        </Grid>
                                    </Grid>
                                </Grid>
                            </div>
                        );
                    }
                    else {
                        return (
                            <ProjectComponent project={this.state.selectedProject} />
                        )
                    }
                }
                else {
                    return (
                        <Grid container className={classes.root} spacing={16}>
                            <div style={{ "marginLeft": "auto", "marginRight": "auto" }}>
                                <h2>Error loading your projects....</h2>
                                <p>{this.state.error}</p>
                            </div>
                        </Grid>
                    )
                }
            } //case admin

            default:
                return null
        }
    }
}








DashboardComponent.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(DashboardComponent);