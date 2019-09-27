import React, { useState, useRef } from 'react';
import classnames from 'classnames';
import { Link } from 'react-router-dom';
import { Location } from 'history';

import { focusTextarea } from 'web/libs/dom';
import { getHomePath } from 'web/libs/paths';
import BooksSelector from './BookSelector';
import { useDispatch, useSelector } from '../../../store';
import { flushContent, markDirty } from '../../../store/editor';
import Textarea from './Textarea';
import Preview from './Preview';
import Button from '../Button';
import styles from './Editor.scss';

interface Props {
  onSubmit: (param: { draftContent: string; draftBookUUID: string }) => void;
  isBusy: boolean;
  cancelPath?: Location<any>;
  isNew?: boolean;
  disabled?: boolean;
  textareaRef: React.MutableRefObject<any>;
  bookSelectorTriggerRef?: React.MutableRefObject<HTMLElement>;
}

enum Mode {
  write,
  preview
}

const Editor: React.SFC<Props> = ({
  onSubmit,
  isBusy,
  disabled,
  textareaRef,
  isNew,
  bookSelectorTriggerRef,
  cancelPath = getHomePath()
}) => {
  const { editor, books } = useSelector(state => {
    return {
      editor: state.editor,
      books: state.books
    };
  });
  const dispatch = useDispatch();
  const [bookSelectorOpen, setBookSelectorOpen] = useState(false);

  const [content, setContent] = useState(editor.content);
  const [mode, setMode] = useState(Mode.write);
  const inputTimerRef = useRef(null);

  const isWriteMode = mode === Mode.write;
  const isPreviewMode = mode === Mode.preview;

  function handleSubmit() {
    // immediately flush the content
    if (inputTimerRef.current) {
      window.clearTimeout(inputTimerRef.current);

      // eslint-disable-next-line no-param-reassign
      inputTimerRef.current = null;
      dispatch(flushContent(content));
    }

    onSubmit({ draftContent: content, draftBookUUID: editor.bookUUID });
  }

  if (disabled) {
    return <div>loading</div>;
  }

  return (
    <form
      id="T-editor"
      className={styles.wrapper}
      onSubmit={e => {
        e.preventDefault();
        handleSubmit();
      }}
    >
      <div className={classnames(styles.row, styles['editor-header'])}>
        <div>
          <BooksSelector
            isReady={books.isFetched}
            isOpen={bookSelectorOpen}
            setIsOpen={setBookSelectorOpen}
            triggerRef={bookSelectorTriggerRef}
            onAfterChange={() => {
              dispatch(markDirty());

              if (textareaRef.current) {
                focusTextarea(textareaRef.current);
              }
            }}
          />
        </div>

        <nav className={styles.tabs}>
          <button
            type="button"
            role="tab"
            aria-selected={isWriteMode}
            className={classnames('button-no-ui', styles.tab, {
              [styles['tab-active']]: isWriteMode
            })}
            onClick={() => {
              setMode(Mode.write);
            }}
          >
            Write
          </button>

          <button
            type="button"
            role="tab"
            aria-selected={isPreviewMode}
            className={classnames('button-no-ui', styles.tab, {
              [styles['tab-active']]: isPreviewMode
            })}
            onClick={() => {
              setMode(Mode.preview);
            }}
          >
            Preview
          </button>
        </nav>
      </div>

      <div className={styles['content-wrapper']}>
        {mode === Mode.write ? (
          <Textarea
            textareaRef={textareaRef}
            inputTimerRef={inputTimerRef}
            content={content}
            onChange={setContent}
            onSubmit={handleSubmit}
          />
        ) : (
          <Preview content={content} />
        )}
      </div>

      <div className={styles.actions}>
        <Button
          id="T-save-button"
          type="submit"
          kind="third"
          size="normal"
          disabled={isBusy}
        >
          {isNew ? 'Save' : 'Update'}
        </Button>

        <Link to={cancelPath} className="button button-second button-normal">
          Cancel
        </Link>
      </div>
    </form>
  );
};

export default Editor;
