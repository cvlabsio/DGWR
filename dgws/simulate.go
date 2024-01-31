package dgws

import (
	"github.com/bcdannyboy/montecargo/dgws/simulator"
	"github.com/bcdannyboy/montecargo/dgws/types"
	"github.com/bcdannyboy/montecargo/dgws/utils"
)

func SimulateIndependentevents(Events []*utils.FilteredEvent, iterations int) ([]*types.SimulationResults, error) {
	SimulationResults := []*types.SimulationResults{}

	for _, event := range Events {
		for i := 0; i < iterations; i++ {
			// probability
			if event.Event.AssociatedProbability == nil {
				if event.Event.AssociatedRisk != nil {
					// associated risk probability
					if event.Event.AssociatedRisk.Probability != nil {
						if event.Event.AssociatedRisk.Probability.SingleNumber != nil {
							// associated risk single number probability
							associatedRiskProbability, err := simulator.SimulateIndependentSingleNumer(event.Event.AssociatedRisk.Probability.SingleNumber)
							if err != nil {
								return nil, err
							}

							SimulationResults = append(SimulationResults, &types.SimulationResults{
								EventID:     event.ID,
								Probability: *associatedRiskProbability,
							})

						} else if event.Event.AssociatedRisk.Probability.Range != nil {
							// associated risk range probability
							associatedRiskProbability, err := simulator.SimulateIndependentRange(event.Event.AssociatedRisk.Probability.Range)
							if err != nil {
								return nil, err
							}

							SimulationResults = append(SimulationResults, &types.SimulationResults{
								EventID:     event.ID,
								Probability: *associatedRiskProbability,
							})

						} else if event.Event.AssociatedRisk.Probability.Decomposed != nil {
							// associated risk decomposed probability
							associatedRiskProbability, associatedRiskProbabilityStandardDeviation, _, _, _, _, err := simulator.SimulateIndependentDecomposed(event.Event.AssociatedRisk.Probability.Decomposed)
							if err != nil {
								return nil, err
							}

							SimulationResults = append(SimulationResults, &types.SimulationResults{
								EventID:                      event.ID,
								Probability:                  associatedRiskProbability,
								ProbabilityStandardDeviation: associatedRiskProbabilityStandardDeviation,
							})
						}
					}
				} else if event.Event.AssociatedMitigation != nil {
					// associated mitigation probability
					if event.Event.AssociatedMitigation.Probability != nil {
						if event.Event.AssociatedMitigation.Probability.SingleNumber != nil {
							// associated mitigation single number probability
							associatedMitigationProbability, err := simulator.SimulateIndependentSingleNumer(event.Event.AssociatedMitigation.Probability.SingleNumber)
							if err != nil {
								return nil, err
							}

							SimulationResults = append(SimulationResults, &types.SimulationResults{
								EventID:     event.ID,
								Probability: *associatedMitigationProbability,
							})
						} else if event.Event.AssociatedMitigation.Probability.Range != nil {
							// associated mitigation range probability
							associatedMitigationProbability, err := simulator.SimulateIndependentRange(event.Event.AssociatedMitigation.Probability.Range)
							if err != nil {
								return nil, err
							}

							SimulationResults = append(SimulationResults, &types.SimulationResults{
								EventID:     event.ID,
								Probability: *associatedMitigationProbability,
							})
						} else if event.Event.AssociatedMitigation.Probability.Decomposed != nil {
							// associated mitigation decomposed probability
							associatedMitigationProbability, associatedMitigationProbabilityStandardDeviation, _, _, _, _, err := simulator.SimulateIndependentDecomposed(event.Event.AssociatedMitigation.Probability.Decomposed)
							if err != nil {
								return nil, err
							}

							SimulationResults = append(SimulationResults, &types.SimulationResults{
								EventID:                      event.ID,
								Probability:                  associatedMitigationProbability,
								ProbabilityStandardDeviation: associatedMitigationProbabilityStandardDeviation,
							})
						}
					}
				}
			} else {
				// default probability item
				if event.Event.AssociatedProbability.SingleNumber != nil {
					// associated probability single number probability
					associatedProbability, err := simulator.SimulateIndependentSingleNumer(event.Event.AssociatedProbability.SingleNumber)
					if err != nil {
						// TODO: handle error
						continue
					}

					SimulationResults = append(SimulationResults, &types.SimulationResults{
						EventID:     event.ID,
						Probability: *associatedProbability,
					})
				} else if event.Event.AssociatedProbability.Range != nil {
					// associated probability range probability
					associatedProbability, err := simulator.SimulateIndependentRange(event.Event.AssociatedProbability.Range)
					if err != nil {
						// TODO: handle error
						continue
					}

					SimulationResults = append(SimulationResults, &types.SimulationResults{
						EventID:     event.ID,
						Probability: *associatedProbability,
					})
				} else if event.Event.AssociatedProbability.Decomposed != nil {
					// associated probability decomposed probability
					associatedProbability, associatedProbabilityStandardDeviation, _, _, _, _, err := simulator.SimulateIndependentDecomposed(event.Event.AssociatedProbability.Decomposed)
					if err != nil {
						// TODO: handle error
						continue
					}

					SimulationResults = append(SimulationResults, &types.SimulationResults{
						EventID:                      event.ID,
						Probability:                  associatedProbability,
						ProbabilityStandardDeviation: associatedProbabilityStandardDeviation,
					})
				}
			}

			// impact
			if event.Event.AssociatedImpact == nil {
				if event.Event.AssociatedRisk != nil {
					if event.Event.AssociatedRisk.Impact != nil {
						// associated risk impact
						if event.Event.AssociatedRisk.Impact.SingleNumber != nil {
							// associated risk single number impact
							associatedRiskImpact, err := simulator.SimulateIndependentSingleNumer(event.Event.AssociatedRisk.Impact.SingleNumber)
							if err != nil {
								return nil, err
							}

							SimulationResults = append(SimulationResults, &types.SimulationResults{
								EventID:      event.ID,
								Impact:       *associatedRiskImpact,
								IsCostSaving: event.Event.AssociatedRisk.Impact.IsCostSaving,
							})
						} else if event.Event.AssociatedRisk.Impact.Range != nil {
							// associated risk range impact
							associatedImpact, err := simulator.SimulateIndependentRange(event.Event.AssociatedRisk.Impact.Range)
							if err != nil {
								return nil, err
							}

							SimulationResults = append(SimulationResults, &types.SimulationResults{
								EventID:      event.ID,
								Impact:       *associatedImpact,
								IsCostSaving: event.Event.AssociatedRisk.Impact.IsCostSaving,
							})
						} else if event.Event.AssociatedRisk.Impact.Decomposed != nil {
							// associated risk decomposed impact
							associatedImpact, associatedImpactStandardDeviation, _, _, _, _, err := simulator.SimulateIndependentDecomposed(event.Event.AssociatedRisk.Impact.Decomposed)
							if err != nil {
								return nil, err
							}

							SimulationResults = append(SimulationResults, &types.SimulationResults{
								EventID:                 event.ID,
								Impact:                  associatedImpact,
								ImpactStandardDeviation: associatedImpactStandardDeviation,
								IsCostSaving:            event.Event.AssociatedRisk.Impact.IsCostSaving,
							})
						}
					} else if event.Event.AssociatedMitigation != nil {
						if event.Event.AssociatedMitigation.Impact != nil {
							// associated mitigation impact
							if event.Event.AssociatedMitigation.Impact.SingleNumber != nil {
								// associated mitigation single number impact
								associatedMitigationImpact, err := simulator.SimulateIndependentSingleNumer(event.Event.AssociatedMitigation.Impact.SingleNumber)
								if err != nil {
									return nil, err
								}

								SimulationResults = append(SimulationResults, &types.SimulationResults{
									EventID:      event.ID,
									Impact:       *associatedMitigationImpact,
									IsCostSaving: event.Event.AssociatedMitigation.Impact.IsCostSaving,
								})
							} else if event.Event.AssociatedMitigation.Impact.Range != nil {
								// associated mitigation range impact
								associatedMitigationImpact, err := simulator.SimulateIndependentRange(event.Event.AssociatedMitigation.Impact.Range)
								if err != nil {
									return nil, err
								}

								SimulationResults = append(SimulationResults, &types.SimulationResults{
									EventID:      event.ID,
									Impact:       *associatedMitigationImpact,
									IsCostSaving: event.Event.AssociatedMitigation.Impact.IsCostSaving,
								})
							} else if event.Event.AssociatedMitigation.Impact.Decomposed != nil {
								// associated mitigation decomposed impact
								associatedMitigationImpact, associatedMitigationImpactStandardDeviation, _, _, _, _, err := simulator.SimulateIndependentDecomposed(event.Event.AssociatedMitigation.Impact.Decomposed)
								if err != nil {
									return nil, err
								}

								SimulationResults = append(SimulationResults, &types.SimulationResults{
									EventID:                 event.ID,
									Impact:                  associatedMitigationImpact,
									ImpactStandardDeviation: associatedMitigationImpactStandardDeviation,
									IsCostSaving:            event.Event.AssociatedMitigation.Impact.IsCostSaving,
								})
							}
						}
					}
				}
			} else {
				// default impact item
				if event.Event.AssociatedImpact.SingleNumber != nil {
					// associated impact single number impact
					associatedImpact, err := simulator.SimulateIndependentSingleNumer(event.Event.AssociatedImpact.SingleNumber)
					if err != nil {
						return nil, err
					}

					SimulationResults = append(SimulationResults, &types.SimulationResults{
						EventID:      event.ID,
						Impact:       *associatedImpact,
						IsCostSaving: event.Event.AssociatedImpact.IsCostSaving,
					})
				} else if event.Event.AssociatedImpact.Range != nil {
					// associated impact range impact
					associatedImpact, err := simulator.SimulateIndependentRange(event.Event.AssociatedImpact.Range)
					if err != nil {
						return nil, err
					}

					SimulationResults = append(SimulationResults, &types.SimulationResults{
						EventID:      event.ID,
						Impact:       *associatedImpact,
						IsCostSaving: event.Event.AssociatedImpact.IsCostSaving,
					})
				} else if event.Event.AssociatedImpact.Decomposed != nil {
					// associated impact decomposed impact
					associatedImpact, associatedImpactStandardDeviation, _, _, _, _, err := simulator.SimulateIndependentDecomposed(event.Event.AssociatedImpact.Decomposed)
					if err != nil {
						return nil, err
					}

					SimulationResults = append(SimulationResults, &types.SimulationResults{
						EventID:                 event.ID,
						Impact:                  associatedImpact,
						ImpactStandardDeviation: associatedImpactStandardDeviation,
						IsCostSaving:            event.Event.AssociatedImpact.IsCostSaving,
					})
				}
			}

			// cost
			if event.Event.AssociatedCost == nil {
				if event.Event.AssociatedMitigation != nil {
					if event.Event.AssociatedMitigation.AssociatedCost != nil {
						// associated mitigation cost
						if event.Event.AssociatedMitigation.AssociatedCost.SingleNumber != nil {
							// associated mitigation single number cost
							associatedMitigationCost, err := simulator.SimulateIndependentSingleNumer(event.Event.AssociatedMitigation.AssociatedCost.SingleNumber)
							if err != nil {
								return nil, err
							}

							SimulationResults = append(SimulationResults, &types.SimulationResults{
								EventID: event.ID,
								Cost:    *associatedMitigationCost,
							})
						} else if event.Event.AssociatedMitigation.AssociatedCost.Range != nil {
							// associated mitigation range cost
							associatedMitigationRangeCost, err := simulator.SimulateIndependentRange(event.Event.AssociatedMitigation.AssociatedCost.Range)
							if err != nil {
								return nil, err
							}

							SimulationResults = append(SimulationResults, &types.SimulationResults{
								EventID: event.ID,
								Cost:    *associatedMitigationRangeCost,
							})
						} else if event.Event.AssociatedMitigation.AssociatedCost.Decomposed != nil {
							// associated mitigation decomposed cost
							associatedMitigationCost, associatedMitigationCostStandardDeviation, _, _, _, _, err := simulator.SimulateIndependentDecomposed(event.Event.AssociatedMitigation.AssociatedCost.Decomposed)
							if err != nil {
								return nil, err
							}

							SimulationResults = append(SimulationResults, &types.SimulationResults{
								EventID:               event.ID,
								Cost:                  associatedMitigationCost,
								CostStandardDeviation: associatedMitigationCostStandardDeviation,
							})
						}
					}
				}
			} else {
				// default cost item
				if event.Event.AssociatedCost.SingleNumber != nil {
					// associated cost single number cost
					associatedCost, err := simulator.SimulateIndependentSingleNumer(event.Event.AssociatedCost.SingleNumber)
					if err != nil {
						return nil, err
					}

					SimulationResults = append(SimulationResults, &types.SimulationResults{
						EventID: event.ID,
						Cost:    *associatedCost,
					})
				} else if event.Event.AssociatedCost.Range != nil {
					// associated cost range cost
					associatedCost, err := simulator.SimulateIndependentRange(event.Event.AssociatedCost.Range)
					if err != nil {
						return nil, err
					}

					SimulationResults = append(SimulationResults, &types.SimulationResults{
						EventID: event.ID,
						Cost:    *associatedCost,
					})
				} else if event.Event.AssociatedCost.Decomposed != nil {
					// associated cost decomposed cost
					associatedCost, associatedCostStandardDeviation, _, _, _, _, err := simulator.SimulateIndependentDecomposed(event.Event.AssociatedCost.Decomposed)
					if err != nil {
						return nil, err
					}

					SimulationResults = append(SimulationResults, &types.SimulationResults{
						EventID:               event.ID,
						Cost:                  associatedCost,
						CostStandardDeviation: associatedCostStandardDeviation,
					})
				}
			}

		}
	}

	return SimulationResults, nil
}
