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
    block: {
        display: "block",
    },
};

class LoginFormComponent extends Component {
    constructor(props) {
        super(props);
        this.state = {
            username: '',
            password: ''
        };
    };
    handleChange = (input) => event => {
        this.setState({ [input]: event.target.value })
    };
    handleSubmit = (event) => {
        //Make a network call somewhere
        console.log(this.state);
        fetch('/api/userAdministration/users')
            .then(data => data.json())
            .then(res => {
                console.log("test")
                console.dir(res);
                 }
            );

        event.preventDefault();
    };

    render() {
        const { classes } = this.props;
        return (
            <form onSubmit={this.handleSubmit}>

                <h2>Login</h2>
                <TextField
                    id="login-username"
                    label="Username"
                    placeholder="Username"
                    className={classes.textField}
                    margin="normal"
                    variant="outlined"
                    onChange={this.handleChange('username')}
                />
                <br></br>

                <TextField
                    id="login-password"
                    label="Password"
                    className={classes.textField}
                    type="password"
                    autoComplete="current-password"
                    margin="normal"
                    variant="outlined"
                    onChange={this.handleChange('password')}
                />
                <br></br>

                <Button variant="outlined" className={classes.button} label="submit" type="submit">
                    Login
                </Button>
            </form>
        )
    }
}

LoginFormComponent.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(LoginFormComponent);