import './LoginScreen.scss';

import React, { ChangeEvent, useState } from 'react';


const LoginScreen = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [agreeTos, setAgreeTos] = useState(false);

    return (
        <div className="login-screen__container">
            <div className="login-screen__login-card">
                <b>Login</b>
                <input 
                    className="" 
                    value={username} 
                    onChange={(e: ChangeEvent<HTMLInputElement>) => setUsername(e.target.value)}
                />
                <input 
                    className="" 
                    value={password} 
                    onChange={(e: ChangeEvent<HTMLInputElement>) => setPassword(e.target.value)}
                />
                 <input
                    type="checkbox"
                    className="" 
                    checked={agreeTos} 
                    onChange={(e: ChangeEvent<HTMLInputElement>) => {
                        setAgreeTos(e.target.checked)
                    } }
                />
            </div>
        </div>
    )
}

export default LoginScreen;