/* Copyright (C) 2019 Monomax Software Pty Ltd
 *
 * This file is part of Dnote.
 *
 * Dnote is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Dnote is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Dnote.  If not, see <https://www.gnu.org/licenses/>.
 */

import React, { useEffect } from 'react';
import { withRouter, RouteComponentProps } from 'react-router-dom';
import NoteGroupList from './NoteGroup/List';
import HeadData from './HeadData';
import { useDispatch, useSelector } from '../../store';
import {
  getFiltersFromSearchStr,
  Filters,
  checkFilterEqual
} from '../../libs/filters';
import { getNotes } from '../../store/notes';
import { groupNotes } from '../../libs/notes';
import { usePrevious } from '../../libs/hooks';
import TopActions from './Actions/Top';
import Flash from '../Common/Flash';

interface Props extends RouteComponentProps {}

function useFetchNotes(filters: Filters) {
  const dispatch = useDispatch();
  const { user } = useSelector(state => {
    return {
      user: state.auth.user.data,
      notes: state.notes
    };
  });
  const prevFilters = usePrevious(filters);

  useEffect(() => {
    if (!user.pro) {
      return () => null;
    }
    if (prevFilters && checkFilterEqual(filters, prevFilters)) {
      return () => null;
    }

    dispatch(getNotes(filters));

    return () => null;
  }, [dispatch, filters, prevFilters, user]);
}

const Home: React.SFC<Props> = ({ location }) => {
  const { notes, user } = useSelector(state => {
    return {
      user: state.auth.user.data,
      notes: state.notes
    };
  });

  const filters = getFiltersFromSearchStr(location.search);
  useFetchNotes(filters);

  const groups = groupNotes(notes.data);

  return (
    <div className="container mobile-nopadding">
      <HeadData filters={filters} />

      <h1 className="sr-only">Notes</h1>

      <Flash kind="danger" when={Boolean(notes.errorMessage)}>
        Error getting notes: {notes.errorMessage}
      </Flash>

      <TopActions />

      <NoteGroupList
        groups={groups}
        pro={user.pro}
        filters={filters}
        isFetched={notes.isFetched}
      />

      {notes.data.length > 10 && <TopActions position="bottom" />}
    </div>
  );
};

export default withRouter(Home);
