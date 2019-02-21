import React, { Component } from 'react';

import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';

import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import ButtonBase from '@material-ui/core/ButtonBase';

import Breadcrumbs from '@material-ui/lab/Breadcrumbs';
import Link from '@material-ui/core/Link';
import FileComponent from './file'

import Icon from '@material-ui/core/Icon'

import Dialog from '@material-ui/core/Dialog';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';


import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import ListSubheader from '@material-ui/core/ListSubheader';

const styles = theme => ({
    root: {
        flexGrow: 1,
    },
    listSection: {
        backgroundColor: 'inherit',
    },
    ul: {
        backgroundColor: 'inherit',
        padding: 0,
    },

    listroot: {
        width: '100%',
        maxWidth: 360,
        backgroundColor: theme.palette.background.paper,
        position: 'relative',
        overflow: 'auto',
        maxHeight: 300,
    },
    title: {
        minWidth: 350,
    },
    paper: {
        height: 180,
        width: 200,
    },
    breadcrumbs: {
        justifyContent: 'center',
        flexWrap: 'wrap',
    },
    crumb: {
        padding: `${theme.spacing.unit}px ${theme.spacing.unit * 2}px`,
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

class ProjectComponent extends Component {
    constructor(props) {
        super(props);
        this.state = {
            username: props.username,
            spacing: '16',
            error: null,
            formError: null,
            project: props.project,
            selectedFile: null,
            newFile: false,


            tag: '',

            //new file
            filename: '',
            tags: []
        };

        console.log(this.state.username)

        //set empty arrays rather than null
        if (this.state.project.files === null)
            this.state.project.files = [];
        if (this.state.project.teams === null)
            this.state.project.teams = [];
        if (this.state.project.users === null)
            this.state.project.users = [];
    }

    componentDidMount() {
        //reloads project data upon mount
        fetch('/api/projectAdministration/project?projectname=' + this.state.project.Projectname, { method: 'GET' })
            .then(data => data.json())
            .then(res => {
                if (res.Success) {
                    if (res.Data == null)
                        this.setState({ project: null })
                    else{
                        
                        if (res.Data[0].files === null)
                            res.Data[0].files = [];
                        if (res.Data[0].teams === null)
                            res.Data[0].teams = [];
                        if (res.Data[0].users === null)
                            res.Data[0].users = [];

                        this.setState({ project: res.Data[0] });
                    }
                }
                else {
                    this.setState({ error: res.Error });
                }
            }
            );
    }
    

    selectFile(file) {
        this.setState({ selectedFile: file })
    }

    handleChange = name => event => {
        this.setState({ [name]: event.target.value, formError: null });
    };

    handleClose = () => {
        this.setState({ filename: '', tags: [], newFile: false });
    };

    reloadProject = () => {
        console.dir(this.state.project)
        fetch('/api/projectAdministration/project?projectname=' + this.state.project.Projectname, { method: 'GET'})
            .then(data => data.json())
            .then(res => {
                if (!res.Success) {
                    alert(res.Error)
                }
                else {
                    if (res.Data == null)
                        this.setState({ filename: '', tags: [], newFile: false, project: [] });
                    else {
                        this.setState({ filename: '', tags: [], newFile: false, project: res.Data[0]});
                    }
                }
            }
            );
    }

    deleteProject = () => {
        console.dir(this.state.project)
        fetch('/api/projectAdministration/project?projectname=' + this.state.project.Projectname, { method: 'DELETE' })
            .then(data => data.json())
            .then(res => {
                if (!res.Success) {
                    alert(res.Error)
                }
                else {
                    //send to dashboard
                    window.location.reload()
                }
            }
            );
    }

    submitFile = () => {
        if (this.state.filename !== ''){
            var newFile = {}
            newFile.filename = this.state.filename
            newFile.versions = [{lastsaved: "20/02/2019", lasteditor: {Username: this.props.username, Password: "s", AccessLevel: "s"}, totaleditTime: "123213", tags: this.state.tags}]
            console.dir(newFile)
            //create new file and reload
            fetch('/api/projectAdministration/fileToProject?projectname=' + this.state.project.Projectname, { method: 'PUT', body: JSON.stringify(newFile) })
                .then(data => data.json())
                .then(res => {
                    if (!res.Success) {
                        alert(res.Error)
                    }
                    else {
                        this.reloadProject()
                    }
                }
                );
        }
        else {
            this.setState({formError: true})
        }
    }

    newFile() {
        this.setState({ newFile: true })
    }

    addTag = () => {
        var newtags = this.state.tags;
        newtags.push(this.state.tag)
        this.setState({ tags: newtags })
    }

    render() {
        const { classes } = this.props;
        const { spacing } = this.state;

        if (this.state.error === null) {
            if (this.state.selectedFile === null) {
                return (
                    <div>
                        <div className={classes.breadcrumbs}>
                            <Paper className={classes.crumb}>
                                <Breadcrumbs arial-label="Breadcrumb">
                                    <Link color="inherit" href="/dashboard" onClick={this.handleClick}>
                                        dashboard
                                    </Link>
                                    <Typography color="inherit">projects</Typography>
                                </Breadcrumbs>
                            </Paper>
                            <br />
                        </div>
                        <Grid id="projectGrid" container className={classes.root} spacing={16}>
                            <Button variant="contained" onClick={() => this.deleteProject()} color="secondary" className={classes.button}>
                                Delete project
                            </Button>
                            <div style={{ "marginLeft": "auto", "marginRight": "auto" }}>
                                <h2>files in {this.state.project.Projectname}</h2>
                            </div>
                            <Grid item xs={12}>
                                <Grid container className={classes.demo} justify="center" spacing={Number(spacing)}>
                                    <Grid key="newFile" item>
                                        <ButtonBase onClick={() => this.newFile()} className={classes.btn}>
                                            <Paper className={classes.paper}>
                                                <Typography gutterBottom variant="display1">
                                                    New file
                                                        </Typography>
                                                <div>
                                                    <Icon className={classes.icon} color="disabled" fontSize="large">
                                                        add_circle
                                                            </Icon>
                                                </div>
                                            </Paper>
                                        </ButtonBase>
                                        <Dialog
                                            open={this.state.newFile}
                                            onClose={this.handleClose}
                                            aria-labelledby="form-dialog-title"
                                            fullwidth = {true}
                                            maxWidth = 'md'
                                        >
                                            <DialogTitle className={classes.title} id="form-dialog-title">New File </DialogTitle>
                                            <DialogContent>

                                                <div>
                                                    <TextField
                                                        error={this.state.formError}
                                                        autoFocus
                                                        margin="dense"
                                                        id="filename"
                                                        label="filename"
                                                        type="text"
                                                        fullwidth="true"
                                                        onChange={this.handleChange('filename')}
                                                    />
                                                    <br></br>
                                                    <List className={classes.root} subheader={<li />}>
                                                        {['tags'].map(sectionId => (
                                                            <li key={`section-${sectionId}`} className={classes.listSection}>
                                                                <ul className={classes.ul}>
                                                                    <ListSubheader>{`${sectionId}`}</ListSubheader>
                                                                    {[this.state.tags].map(item => (
                                                                        <ListItem  key={`tags:-${sectionId}-${item}`}>
                                                                            <ListItemText primary={`${item}`} />
                                                                        </ListItem>
                                                                    ))}
                                                                </ul>
                                                            </li>
                                                        ))}
                                                    </List>
                                                    <TextField
                                                        autoFocus
                                                        margin="dense"
                                                        id="tag"
                                                        label="tag"
                                                        type="text"
                                                        fullwidth="true"
                                                        onChange={this.handleChange('tag')}
                                                    />
                                                    <Button onClick={this.addTag} color="primary">
                                                        Add
                                                    </Button>
                                                </div>
                                            </DialogContent>
                                            <DialogActions>
                                                <Button onClick={this.handleClose} color="primary">
                                                    Cancel
                                                    </Button>
                                                <Button onClick={this.submitFile} color="primary">
                                                    Create
                                                    </Button>
                                            </DialogActions>
                                        </Dialog>
                                    </Grid>
                                    {this.state.project.files.map(value => (
                                        <Grid key={value.filename} item>
                                            <ButtonBase onClick={() => this.selectFile(value)} className={classes.btn}>
                                                <Paper className={classes.paper}>
                                                    <Typography gutterBottom variant="display2">
                                                        {value.filename}
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
                    <FileComponent file={this.state.selectedFile} project={this.state.project} username={this.state.username} />
                )
            }
        }
        else {
            return (
                <Grid container className={classes.root} spacing={16}>
                    <div style={{ "margin-left": "auto", "margin-right": "auto" }}>
                        <h2>Error loading project....</h2>
                        <p>{this.state.error}</p>
                    </div>
                </Grid>
            )
        }
    }
}








ProjectComponent.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ProjectComponent);