import React, { Component } from 'react';

import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';

import Grid from '@material-ui/core/Grid';

import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';

import Breadcrumbs from '@material-ui/lab/Breadcrumbs';
import Link from '@material-ui/core/Link';

import ProjectComponent from './project'

const styles = theme => ({
    root: {
        flexGrow: 1,
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
            error: null,
            file: props.file,
            selectedFile: null
        };
    }


    handleClick() {
        this.setState({file: null})
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
                                    <Link style={{"cursor":"pointer"}} color="inherit" onClick={() => this.handleClick()}>
                                        projects
                                    </Link>
                                    <Typography color="inherit">{this.state.file.filename}</Typography>
                                </Breadcrumbs>
                            </Paper>
                            <br />
                        </div>
                        <p>File view shall go here...</p>
                    </div>
                );
            }
            else {
                return <ProjectComponent project={this.props.project} />
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