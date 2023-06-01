import { Visibility, VisibilityOff } from '@mui/icons-material'
import { FormControl, IconButton, InputAdornment, InputLabel, OutlinedInput } from '@mui/material'
import React from 'react'

const PasswordInput = ({ values, setValues }) => {
  const [showPassword, setShowPassword] = React.useState(false);

  const handleClickShowPassword = () => setShowPassword((show) => !show);

  const handleMouseDownPassword = (event) => {
    event.preventDefault();
  };
  return (
    <FormControl sx={{ width: '100%' }} variant="outlined">
      <InputLabel size="small" htmlFor="outlined-adornment-password">Contraseña</InputLabel>
      <OutlinedInput
        size="small"
        id="outlined-adornment-password"
        type={showPassword ? 'text' : 'password'}
        value={values.password}
        onChange={(e) => setValues({ ...values, password: e.target.value })}
        endAdornment={
          <InputAdornment position="end">
            <IconButton
              aria-label="toggle password visibility"
              onClick={handleClickShowPassword}
              onMouseDown={handleMouseDownPassword}
              edge="end"
            >
              {showPassword ? <VisibilityOff /> : <Visibility />}
            </IconButton>
          </InputAdornment>
        }
        label="Contraseña"
      />
    </FormControl>
  )
}

export default PasswordInput