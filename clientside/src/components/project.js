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

class ProjectComponent extends Component {
    constructor(props) {
        super(props);
        this.state = {
            spacing: '16',
            error: null,
            project: props.project,
            selectedFile: null
        };
    }

    selectFile(file) {
        this.setState({selectedFile: file})
    }

    render() {
        const { classes } = this.props;
        const { spacing } = this.state;

        if (this.state.error === null) {
            if (this.state.selectedFile === null){
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
                    <div style={{ "marginLeft": "auto", "marginRight": "auto" }}>
                        <h2>files in {this.state.project.Projectname}</h2>
                    </div>
                    <Grid item xs={12}>
                        <Grid container className={classes.demo} justify="center" spacing={Number(spacing)}>
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
                    <FileComponent file={this.state.selectedFile} project={this.state.project} />
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