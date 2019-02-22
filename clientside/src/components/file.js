import React, { Component } from 'react';

import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';

import Grid from '@material-ui/core/Grid';

import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';

import Breadcrumbs from '@material-ui/lab/Breadcrumbs';
import Link from '@material-ui/core/Link';

import ProjectComponent from './project'

import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

import Dialog from '@material-ui/core/Dialog';
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
        width: '50%'
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

class FileComponent extends Component {
    constructor(props) {
        super(props);
        this.state = {
            username: props.username,
            error: null,
            file: props.file,
            initialLength: 0,
            project: props.project,
            selectedFile: null,



            newVersion: false,
            tag: '',
            tags: [],
            formError: null
        };

        if (this.state.file.versions === null)
            this.state.file.versions = [];


        this.state.file.versions = this.state.file.versions.reverse()
        this.state.tags = this.state.file.versions[0].tags

        if (this.state.tags === null) {
            this.state.tags = []
            this.state.initialLength = 0
        } else {
            this.state.initialLength = this.state.tags.length
        }

        console.log(this.state.username)
    }


    handleClick() {
        this.setState({ file: null })
    }

    deleteFile = () => {

        var body = {
            projectname: this.state.project.Projectname,
            files: [this.state.file]
        }


        fetch('/api/projectAdministration/removeFiles', { method: 'PUT', body: JSON.stringify(body) })
            .then(data => data.json())
            .then(res => {
                if (!res.Success) {
                    alert(res.Error)
                }
                else {
                    //send to dashboard
                    this.setState({ file: null })
                }
            }
            );
    }

    editFile() {
        this.lockFile();
        this.setState({ newVersion: true })
    }

    lockFile() {
        //would hit db and lock the file, anyone trying to edit would be told they can't
    }

    unlockFile() {

    }

    submitVersion = () => {
        if (this.state.tags.length > this.state.initialLength) {

            if (typeof this.state.username !== 'undefined'){
            var newVersion = {
                filename: this.state.file.filename,
                version: {
                    lastsaved: this.state.file.versions[0].lastsaved + "1",
                    lasteditor: this.state.username,
                    totaleditTime: this.state.file.versions[0].totaleditTime + "2",
                    tags: this.state.tags
                }
            }
        }else {
                var newVersion = {
                    filename: this.state.file.filename,
                    version: {
                        lastsaved: this.state.file.versions[0].lastsaved + "1",
                        lasteditor: "",
                        totaleditTime: this.state.file.versions[0].totaleditTime + "2",
                        tags: this.state.tags
                    }
                }
        }
            console.dir(newVersion)
            // add file version
            fetch('/api/fileAdministration/addFileVersion', { method: 'PUT', body: JSON.stringify(newVersion) })
                .then(data => data.json())
                .then(res => {
                    if (!res.Success) {
                        alert(res.Error)
                    }
                    else {
                        this.reloadFile()
                    }
                }
                );
        }
        else {
            this.setState({ formError: true })
        }
    }

    reloadFile = () => {
        fetch('/api/fileAdministration/file?filename=' + this.state.file.filename, { method: 'GET' })
            .then(data => data.json())
            .then(res => {
                if (!res.Success) {
                    alert(res.Error)
                }
                else {
                    if (res.Data == null)
                        this.setState({ newVersion: false, file: { filename: this.state.file.filename, versions: [] }, tags: [] });
                    else {
                        res.Data[0].versions = res.Data[0].versions.reverse()
                        this.setState({ newVersion: false, file: res.Data[0], tags: res.Data[0].versions[0].tags });
                    }
                }
            }
            );
    }

    handleChange = name => event => {
        this.setState({ [name]: event.target.value, formError: null });
    };

    handleClose = () => {
        this.setState({ newVersion: false });
        this.unlockFile();
    };

    addTag = () => {
        var newtags = this.state.tags;
        newtags.push(this.state.tag)
        this.setState({ tags: newtags })
    }


    formatTags(tags) {
        var tagstr = ""

        if (tags !== null) {
            for (var i = 0; i < tags.length; i++) {
                tagstr += tags[i] + ", "
            }
        }

        return tagstr
    }

    render() {
        const { classes } = this.props;
        // const { spacing } = this.state;

        if (this.state.error === null) {
            if (this.state.file !== null) {
                return (
                    <div>
                        <div className={classes.breadcrumbs}>
                            <Paper className={classes.crumb}>
                                <Breadcrumbs arial-label="Breadcrumb">
                                    <Link color="inherit" href="/dashboard">
                                        dashboard
                                    </Link>
                                    <Link style={{ "cursor": "pointer" }} color="inherit" onClick={() => this.handleClick()}>
                                        projects
                                    </Link>
                                    <Typography color="inherit">{this.state.file.filename}</Typography>
                                </Breadcrumbs>
                            </Paper>
                            <br />
                        </div>
                        <Button style={{ "width": "auto" }} variant="contained" onClick={() => this.editFile()} color="primary" className={classes.button}>
                            Edit file
                        </Button>
                        <Dialog
                            open={this.state.newVersion}
                            onClose={this.handleClose}
                            aria-labelledby="form-dialog-title"
                            fullwidth='true'
                            maxWidth='md'
                        >
                            <DialogTitle className={classes.title} id="form-dialog-title">Edit File </DialogTitle>
                            <DialogContent>

                                <div>
                                    <List className={classes.root} subheader={<li />}>
                                        {['tags'].map(sectionId => (
                                            <li key={`section-${sectionId}`} className={classes.listSection}>
                                                <ul className={classes.ul}>
                                                    <ListSubheader>{`${sectionId}`}</ListSubheader>
                                                    {[this.state.tags].map(item => (
                                                        <ListItem key={`tags:-${sectionId}-${item}`}>
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
                                <Button onClick={this.submitVersion} color="primary">
                                    Edit
                                                    </Button>
                            </DialogActions>
                        </Dialog>
                        <div style={{ "marginLeft": "auto", "marginRight": "auto" }}>
                            <h2>{this.state.file.filename}</h2>
                        </div>
                        <div className={classes.root} style={{ "marginLeft": "auto", "marginRight": "auto" }}>
                            {this.state.file.versions.map(version => (
                                <ExpansionPanel key={this.state.file.versions.indexOf(version)}>
                                    <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
                                        <Typography className={classes.heading}>{version.lasteditor} - {version.lastsaved}</Typography>
                                    </ExpansionPanelSummary>
                                    <ExpansionPanelDetails>
                                        <Typography color="inherit">
                                            total edit time: {version.totaleditTime}
                                            <br></br>
                                            tags : {this.formatTags(version.tags)}
                                        </Typography>
                                    </ExpansionPanelDetails>
                                </ExpansionPanel>
                            ))}

                        </div>

                        <br></br>
                        <br></br>
                        <Button style={{ "width": "auto" }} variant="contained" onClick={() => this.deleteFile()} color="secondary" className={classes.button}>
                            Remove file from project
                        </Button>

                    </div>
                );
            }
            else {
                return <ProjectComponent project={this.state.project} />
            }
        }
        else {
            return (
                <Grid container className={classes.root} spacing={16}>
                    <div style={{ "margin-left": "auto", "margin-right": "auto" }}>
                        <h2>Error loading file....</h2>
                        <p>{this.state.error}</p>
                    </div>
                </Grid>
            )
        }
    }
}








FileComponent.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(FileComponent);