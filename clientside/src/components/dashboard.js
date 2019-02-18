import React, { Component } from 'react';

import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';

const styles = {
    root: {
        flexGrow: 1,
    },
    grow: {
        flexGrow: 1,
    },
    block: {
        display: "block",
    },
};

class DashboardComponent extends Component {
    constructor(props) {
        super(props);
        this.state = {
            username: '',
            password: ''
        };
    };
    

    render() {
        const { classes } = this.props;
        return(
            <p>hey llol</p>
        )
    }
}

DashboardComponent.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(DashboardComponent);