import React, { useEffect } from 'react'
import { connect } from 'react-redux'
import { aggregationOperations } from 'state/ducks/aggregation'
import { NetworkGraph } from 'components/networkGraph'
import { NetworkGraphInfo } from 'components/networkGraphInfo'
import { AnswerHistory } from 'components/answerHistory'
import { DeviationHistory } from 'components/deviationHistory'
import withMainnet from 'enhancers/withMainnet'
import { OracleTable } from 'components/oracleTable'

const NetworkPage = ({ initContract, clearState, options }) => {
  useEffect(() => {
    /* eslint-disable-next-line @typescript-eslint/no-empty-function */
    initContract(options).catch(() => {})
    return () => {
      clearState()
    }
  }, [initContract, clearState, options])

  return (
    <div className="page-wrapper network-page">
      <NetworkGraph options={options} />
      <NetworkGraphInfo options={options} />
      {options.history && <AnswerHistory options={options} />}
      {options.history && <DeviationHistory options={options} />}
      <OracleTable />
    </div>
  )
}

const mapDispatchToProps = {
  initContract: aggregationOperations.initContract,
  clearState: aggregationOperations.clearState,
}

export default connect(null, mapDispatchToProps)(withMainnet(NetworkPage))
