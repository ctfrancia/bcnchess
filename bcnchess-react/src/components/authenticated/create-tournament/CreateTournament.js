import React, { useState } from 'react';
import { useSelector } from 'react-redux';
import CTReview from './CTReview';
import CTForm from './CTForm';

const CreateTournament = () => {
  const isBeingEdited = useSelector(state => state.createTournamentFormData.isBeingEdited)
  
  const [poster, setPoster] = useState()

  return isBeingEdited ? <CTForm poster={poster} onPosterChange={setPoster} /> : <CTReview />
};

export default CreateTournament;