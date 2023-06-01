import styled, { css } from "styled-components";

export const FormContainer = styled.div`
    padding: 30px 10px;
    display: flex;
    flex-direction: column;
    width: 350px;
    gap: 10px;
    border-radius: 10px;
    border: 2px solid rgb(247, 248, 250);
    background-color: #F2BF5E;
    margin-left: 35%;
`;

export const Container = styled.div`  
    display: flex;
    flex-direction: column;
    text-align: left;
    ${({ content }) => content === "row" && css`
        width: 100%;
        flex-direction: row;
        gap: 10px;
        justify-content: center;
    `}
`;

export const Input = styled.input`
    padding: 5px;
    border-radius: 5px;
    outline: none;
    border: 1px solid #4E6E81;
`;

export const Title = styled.h1`
    text-align: center;
    font-size: 40px;
`;


export const Error = styled.p `
    color: #C70039;
    font-size: "10px"; 
    padding: 1px;
    margin: 1px 0 0 6px; 
    font-style: italic;
    font-family: monospace;
`;
