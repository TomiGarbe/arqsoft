export const isLoggedIn = () => {
    const Token = localStorage.getItem('Token');
    return Token !== null;
};