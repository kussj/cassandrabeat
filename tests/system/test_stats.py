from cassandrabeat import TestCase
#import os


"""
Contains tests for table statistics.
"""


class Test(TestCase):
    def test_table_stats(self):
        """
        Checks that table stats are found in the
        output and have the expected types.
        """
        self.render_config_template()
        cassandrabeat = self.start_cassandrabeat()
        self.wait_until(lambda: self.output_has(lines=1))
        cassandrabeat.kill_and_wait()

        output = self.read_output()[0]

        for key in [
            "table_name",
        ]:
            assert type(output[key].encode('ascii','ignore')) is str

        for key in [
            "read_latency",
            "write_latency",
        ]:
            assert type(output[key]) is float
