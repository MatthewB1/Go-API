import React, { Component } from 'react';


import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button'
// import ArrowBackIosIcon from '@material-ui/icons/ArrowBackIos';
// import ButtonBase from '@material-ui/core/ButtonBase';


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

    logout(){
        localStorage.setItem('token', '')
        const { history } = this.props
        history.push('/login')
    }

    componentWillMount(){
        //no reason to ever go to this url, send user to dashboard
        const { history } = this.props
        history.push('/dashboard')
    }

    render(){
        const { classes } = this.props;
        return (
            <div >
                <AppBar className= {classes.appBar} position="static">
                    <Toolbar>
                        {/* <ButtonBase onClick={() => this.goBack()} className={classes.btn}>
                        <ArrowBackIosIcon className={classes.icon} />
                        </ButtonBase> */}
                        <Typography variant="h6" color="inherit" className={classes.grow}>
                            TMS
                         </Typography>
                        <Button id="logoutButton" onClick={() => this.logout()} ref="logoutButton" color="inherit">Logout</Button>
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