import React, { Component } from 'react';

import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField'

const styles = {
    root: {
        flexGrow: 1,
    },
    grow: {
        flexGrow: 1,
    },
};

class LoginFormComponent extends Component {
    constructor() {
        super();
        this.state = {};
    }
    handleChange = (event) => {
        //
    }
    handleSubmit = (event) => {
        //Make a network call somewhere
        event.preventDefault();
    }
    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <TextField floatingLabelText="username" onChange={this.handleChange} />
                <TextField floatingLabelText="password" onChange={this.handleChange} />
                <Button label="Submit" type="submit" />
            </form>
        )
    }
}

LoginFormComponent.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(LoginFormComponent);