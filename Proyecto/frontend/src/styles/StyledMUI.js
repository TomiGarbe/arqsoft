import { Box } from "@mui/material";
import { styled } from "@mui/material/styles";
import bg from "./teeth.jpg";

export const AuthContainer = styled(Box)`
  background-image: url(${bg});
  background-size: cover;
  min-height: 100%;
  position: fixed;
  width: 100%;
  background-repeat: no-repeat;
  background-color: #4444;
`;

export const Glass = styled(Box)`
  background: rgba(255, 255, 255, 0.17);
  border-radius: 16px;
  box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(5px);
  -webkit-backdrop-filter: blur(5px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  width: 400px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  text-align: center;
  height: auto;
  padding: 20px;
  position: absolute;
  top: 50%;
  left: 35%;
  transform: translate(-50%, -50%);
`;
