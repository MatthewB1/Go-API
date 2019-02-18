import React, { Component } from 'react';

import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button'

const styles = {
    root: {
        flexGrow: 1,
    },
    grow: {
        flexGrow: 1,
    },
    menuButton: {
        marginLeft: -12,
        marginRight: 20,
    },
    appBar: {
        backgroundColor: "#b29600",
    },
};

class MainContentComponent extends Component {

    constructor(props){
        super(props);
        this.state = {};
    }

    render(){
        const { classes } = this.props;
        return (
            <div >
                <AppBar className= {classes.appBar} position="static">
                    <Toolbar>
                        <Typography variant="h6" color="inherit" className={classes.grow}>
                            TMS
                         </Typography>
                        <Button id="logoutButton" ref="logoutButton" color="inherit">Logout</Button>
                    </Toolbar>
                </AppBar>
            </div>
        )
    }
}

MainContentComponent.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(MainContentComponent);