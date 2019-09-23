import React, { useState } from 'react';

import authStyles from '../../Common/Auth.scss';
import Button from '../../Common/Button';

interface Props {
  onResetPassword: (password: string, passwordConfirmation: string) => void;
  submitting: boolean;
}

const PasswordResetConfirmForm: React.SFC<Props> = ({
  onResetPassword,
  submitting
}) => {
  const [password, setPassword] = useState('');
  const [passwordConfirmation, setPasswordConfirmation] = useState('');

  return (
    <form
      onSubmit={e => {
        e.preventDefault();

        onResetPassword(password, passwordConfirmation);
      }}
      className={authStyles.form}
    >
      <div className={authStyles['input-row']}>
        <label htmlFor="password-input" className={authStyles.label}>
          Password
          <input
            id="password-input"
            type="password"
            placeholder="********"
            className="form-control"
            value={password}
            onChange={e => {
              const val = e.target.value;

              setPassword(val);
            }}
          />
        </label>
      </div>

      <div className={authStyles['input-row']}>
        <label
          htmlFor="password-confirmation-input"
          className={authStyles.label}
        >
          Password Confirmation
          <input
            id="password-confirmation-input"
            type="password"
            placeholder="********"
            className="form-control"
            value={passwordConfirmation}
            onChange={e => {
              const val = e.target.value;

              setPasswordConfirmation(val);
            }}
          />
        </label>
      </div>

      <Button
        type="submit"
        size="normal"
        kind="first"
        stretch
        className={authStyles['auth-button']}
        isBusy={submitting}
      >
        Reset password
      </Button>
    </form>
  );
};

export default PasswordResetConfirmForm;
