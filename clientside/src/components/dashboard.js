import React, { Component } from 'react';

import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';

import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import ButtonBase from '@material-ui/core/ButtonBase';

import Breadcrumbs from '@material-ui/lab/Breadcrumbs';

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
});

class DashboardComponent extends Component {
    constructor(props) {
        super(props);
        this.state = {
            spacing: '16',
            username: '',
            accessLevel: '',
            error: null,
            projects: [],
            selectedProject: null
        };
    }

    componentDidMount() {
        const token = decode(localStorage.getItem('token'));
        this.setState({username: token.name, accessLevel : token.admin ?("admin") : ("user")})


        fetch('/api/projectAdministration/projects?user=' + token.name, { method: 'GET' })
            .then(data => data.json())
            .then(res => {
                if (res.Success) {
                    this.setState({ projects: res.Data});
                }
                else {
                    this.setState({ error: res.Error });
                }
            }
            );
    }

    usersProjects(projects){
        var usersprojects = [];
        console.dir(usersprojects)

        for (var project in projects){
            if (project.Projectlead.username === this.state.username){
                usersprojects.push(project);
            }
            for (var team in project.teams){
                if (team.teamleader.username === this.state.username){
                    usersprojects.push(project);
                }
                for (var member in team.teamMembers){
                    if (member.username === this.state.username){
                        usersprojects.push(project);
                    }
                }
            }
            for (var user in project.users){
                if (user.username === this.state.username){
                    usersprojects.push(project);
                }
            }
        }

        usersprojects.push(projects[0])
        console.dir(usersprojects);
        return usersprojects;
    }

    selectProject(value) {
        this.setState({ selectedProject: value })
    }

    handleClick = event => {

    }

    render() {
        const { classes } = this.props;
        const { spacing } = this.state;

        if (this.state.error === null) {
            if (this.state.selectedProject === null) {
                return (
                    <div>
                        <div className={classes.breadcrumbs}>
                            <Paper className={classes.crumb}>
                                <Breadcrumbs arial-label="Breadcrumb">
                                    {/* <Link color="inherit" href="/dashboard" onClick={this.handleClick}>
                                        dashboard
                                    </Link> */}
                                    <Typography color="inherit">dashboard</Typography>
                                </Breadcrumbs>
                            </Paper>
                            <br />
                        </div>
                        <Grid container className={classes.root} spacing={16}>
                            <div style={{ "marginLeft": "auto", "marginRight": "auto" }}>
                                <h2>Hi {this.state.username}!</h2>
                            </div>
                            <Grid item xs={12}>
                                <Grid container className={classes.demo} justify="center" spacing={Number(spacing)}>
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
                    <div style={{ "margin-left": "auto", "margin-right": "auto" }}>
                        <h2>Error loading your projects....</h2>
                        <p>{this.state.error}</p>
                    </div>
                </Grid>
            )
        }
    }
}








DashboardComponent.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(DashboardComponent);